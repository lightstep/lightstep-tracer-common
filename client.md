# LightStep Tracer Client v3 Specification

LightStep Engineering
May 3, 2018

## Overview

This document serves to specify how LightStep client libraries
implement the OpenTracing
[specification](https://github.com/opentracing/specification/blob/master/specification.md).
Like that document, this is intended to coordinate cross-language
development and necessarily stays away from language-specific and
toolchain-specific discussion.

This specification defines the third generation of LightStep client
libraries.  Whereas first generation client libraries offered Thrift
transport, and second generation client libraries added gRPC support,
third generation libraries remove both of these methods in favor of
plain http/2 transport using either protobuf or JSON encoding.  Third
generation client libraries are expected to perform efficiently in
both high-throughput and low-resource environments.  Third generation
client libraries offer a user-defined transport option, allowing
applications to customize transport and augment tracing data outside
the process.

LightStep tracer libraries will be packaged and built such that users
have fine-grain control over which tracer dependencies are used at
runtime, using standard build tools and package-distribution
mechanisms for each language.  This document currently covers efforts
to upgrade client libraries to match this specification across three
server-language platforms (Java/JRE, Golang, and C++11) and two
mobile-language platforms (Java/Android, Obj-C/iOS).  This document
also describes the appropriate level of testing and benchmarking to
include with each client library.

### Motivation

Our motivation for these changes is to improve the breadth of
operating conditions for which LightStep's tracer library can provide
efficient, reliable diagnostics reporting.  These implementations are
already constrained on both sides, with the OpenTracing API
specification and LightStep collector protocol defining the input and
output semantics.

The primary task of this document is to specify how client libraries
should use the available resources efficiently, particularly on the
performance boundary, where we may be forced to drop data or increase
resource usage, depending on service conditions. We identify the most
important scenarios for consideration as follows:

* New spans produced while the buffer of pending data is nearly full
* Non-retryable failures received before the timeout
* Reporting timeout received
* User code creates excessively large spans or logs

In these cases, this document will specify how clients should behave,
in the mutual interest of valuing client resources, delivering
instrumentation reliably, and protecting LightStep collectors from
overload.

## Reporting protocol: user-visible fields

LightStep's [collector
protocol](https://github.com/lightstep/lightstep-tracer-common/blob/master/collector.proto)
defines the data model used for ingesting spans that was adopted
for second-generation tracers, introduced with the migration to
gRPC.  There is a straightforward translation from OpenTracing
concept to members of these fields with little room for
interpretation.  User-visible fields are called out in `collector.proto`.

We document the interpretation of certain fields here to support
deeper-integration between the application and the mode of
user-defined transport.  To that end, we only document those
fields necessary to support approved uses.  

### Report structure

The `lightstep.ReportRequest` message consists of three
user-interpretable fields:

1. **reporter**: describes the client process that generated the batch of spans
2. **auth**:     describes the access token associated with the batch of spans
3. **spans**:    the main payload, a list of span structs

Reports must be smaller than 4MB when serialized as a binary-encoded
protocol message.

#### Reporter

The `reporter` field describes the client process in terms of a
unique uuid (`lightstep.Reporter.reporter_id`) and list of tag
values.  LightStep reserves the use of tag names that begin
`lightstep.` and will specify several reporter tag values by
default, including:

Key | Meaning
----|--------
`lightstep.component_name`          | This maps to "service" in LightStep's UI, not to be confused with the OpenTracing ["component"](https://github.com/opentracing/specification/blob/master/semantic_conventions.md) semantic concept.
`lightstep.hostname`                | LightStep clients set this to the operating system hostname.
`lightstep.tracer_platform`         | A string describing the particular language and runtime (e.g., "iOS")
`lightstep.tracer_platform_version` | A string describing the version of the client library (e.g., "iOS 11.4")
`lightstep.command_line`            | Formatted command-line arguments.
`lightstep.tracer_version`          | Version of the tracer library itself.


Reporter fields are effectively applied to each of the spans
contained within the report.  Spans may override each of the
above values by setting the same tag in their own
`lightstep.Span.tags` set.

#### Auth

The `auth` field contains the client's access token.  This field
is set directly from the value passed to the tracer constructor.
Applications can late-bind the access token used by overriding
this field downstream in their user-defined transport.

#### Spans

The main report payload consists of a list of spans.
`lightstep.Span` matches the OpenTracing `Span` specification
closely.  Users may apply tag values to spans downstream in their
user-defined transport layer, by adding them to
`lightstep.Span.tags`.

Span tags override reporter tags, making it possible to create spans
with a client on behalf of another process.  To set the reporter GUID
for a span, use:
 
Key | Meaning
----|--------
`lightstep.guid` | Equivalent to the reporter uuid.  If 64-bits, use a 16-byte hex string representation.  If 128-bits, use a uuid.v4 string representation.

## Library Design

LightStep client libraries will be factored to include a default
http/2 implementation and a user-defined transport mode.  Although
gRPC is eliminated from third generation client libraries, we use the
same protocol definition and are committed to gRPC support on the
server, meaning that user-defined transport implementations may use
the gRPC [`lightstep.collector.CollectorService.Report`](https://github.com/lightstep/lightstep-tracer-common/blob/4c649d1a7ac52b9cafc7f8d21fe304f8fa4a4ae3/collector.proto#L105) endpoint.

Client libraries may be viewed as an assembly of several parts:
one a pure tracing implementation, one for buffering and flushing
spans, one for encoding span batches, and one for transporting
span batches.  We label these parts:

1. Pure Tracing: this is an implementation of the OpenTracing Tracer interface, which handles translation from OpenTracing API calls into in-memory structures.  This component also implements OpenTracing Inject and Extract operations.
1. Span Recorder: this component receives finished spans from the pure tracing component.  Users may supplied their own span recorder for user-defined transport.  This module is expected not to block the caller. This module implements Flush support and is generally responsible for limiting resource usage.
1. Report Builder: this component contains logic to encode a LightStep report from a set of finished 1.
spans.
1. Transporter: this component is responsible for sending a report batch to LightStep over HTTP.

There are several competing interests present when designing client
libraries, which we prioritize as follows:

1. Do not harm the application
1. Do not harm the LightStep service
1. Do not drop data
1. Do not waste resources

First and second generation LightStep client libraries offered several
resource- and performance-related parameters, notably: (1) limit on
number of buffered spans, (2) minimum flush period, (3) maximum flush
period, (4) report timeout.  Those libraries were restricted to a
maximum of one simultaneous Report being sent at a time, making these
parameters easy to interpret but difficult to tune.  Third generation
client libraries will be configured by a new set of parameters:

Name                   | Interpretation (implemented by)
---------------------- | --------------
`max_memory`           | Number of bytes of memory buffered (Span Recorder)
`max_report_size`      | Limits the size of an outgoing report (Report Builder and/or Span Recorder)
`max_concurrency`      | Number of CPUs dedicated to sending reports (Transporter)
`max_bytes_per_second` | Prevent sending more frequently (Span Recorder)

When user-defined transport is selected, only the `max_report_size` setting
is meaningful.

Clients using built-in transport may presume that the library
consumes up to `max_memory` bytes of buffered spans plus up to
`max_report_size` of memory per concurrent sender.  The
priorities listed above should be used to determine client
behavior, without need for additional tuning parameters.

#### Note about mobile platforms

We expect that mobile platforms will require separate transport
implementations compared with the same language on a non-mobile
platform.

#### Note about built-in safety

Client libraries are expected to use built-in facilities such as
string formatting and JSON marshalling when producing reports, and are
therefore only as safe as those facilities.  This risk is passed on the
programmer.

#### Note protocol about buffers vs. JSON

Protocol buffer library support varies significantly by language,
and in some languages there is more than one viable choice of
library.  The pure tracing implementation should be architected
so as not to constrain the choice of library used for encoding
span data in the report builder and transporter, while also not
constraining performance.  This implies that the Pure Tracing
component use generic programming techniques, to avoid needless
copying of data into the Report Builder; it also implies that we
expect the Report Builder and Transporter to be tightly coupled.

### Factorization

The goal of user-defined transport is to provide alternative
implementations for the Span Recorder and Transporter components,
while still relying on the Pure Tracing component and a Report
Builder.

This factorization already exists in some of the libraries, for
example a `Recorder` in
[C++](https://github.com/lightstep/lightstep-tracer-cpp/blob/4ea8bda9aed08ad45d6db2a030a1464e8d9b783f/src/recorder.h#L9)
and a `SpanRecorder` in
[Golang](https://github.com/lightstep/lightstep-tracer-go/blob/644c3d5ecbd0499c50a1329f89ba287921fc1144/options.go#L66),
but the interface is not currently consistent.  Java has multiple
transport options, but no facility for a user-provided transport,
while Objective-C has only a single transport option.

#### Definitions

There are two kinds of object being described here.

- A "concrete interface object" is the implementation of an OpenTracing interface
- A "data-transport object" is a structured object that can be encoded for LightStep

A tracing library provides concrete interface objects to the
user, while employing data-trasport objects to manage storing and
encoding internal state.

#### Pure Tracing component

The `AbstractTracer` type provides the basic implementation of
`opentracing.Tracer`.  Concrete implementations will:

- Provides external system dependencies (e.g., scope manager, logger, clock, ID generator)
- Factory for concrete Span (or SpanBuilder) and SpanContext interface objects
- Factory for concrete data-transfer objects (timestamps, logs, span references, key values) of in-flight Span state
- Provides the opentracing Inject and Abstract APIs.

The `AbstractSpan` type provides the basic implementation of
`opentracing.Span`.  Concrete implementations will:

- Contains a reference to an `AbstractTracer` implementation
- Contains a reference to its own span recorder
- Contains the underlying Span data-transfer object.

The `SpanContext` type provides a LightStep implementation of
`opentracing.SpanContext`.

#### Span Recorder component

The Span Recorder component is responsible for the logic of
buffering and flushing data, deciding when to drop, whether to
increase or decrease concurrency and report size, and when to
back-off if the service is not responding.  It is cheifly
constrained by the limits on memory, CPU, and network usage.

TODO: Because the Span Recorder component is not required for
deploying a user-defined transport option, detailed deisgn of the
Span Recorder component will be future LightStep engineering
work.  LightStep plans to begin constructing a v3 client for
Golang in June 2018, to flush out this design.

##### Technical note on limiting report size

LightStep places a limit on overall report size for pragmatic
reasons, because it assists with load balancing and because we
use gRPC downstream, which imposes a hard limit.  Limiting report
size is not easy to do cheaply, however, because it takes
significant CPU to precisely calculate the encoded size of a
message.

We have considered several approaches to deal with this:

- A precise algorithm: as reports are built incrementally, re-compute the total size by asking the object its encoded size when each new span is added, stop and remove the last span when it exceeds the limit.  This requires O(N^2) calls to compute the encoded size and O(N) calls to encode a span.
- An approximate algorithm: as reports are built incrementally, compute the estimated size of each span and extrapolate the estimated report size using assumptions about protobuf overhead.  This requires O(N) calls to copmute the encoded size.
- An efficient algorithm: encode each span individually into a byte array, yielding both the encoding and its size.  Assemble protobuf reports by concatenating byte arrays, which works because `lighstep.Report` contains a repeated `lightstep.Span` field at the outermost level.  There is nearly zero overhead to this approach, though it couples the Span Recorder to the Report Builder to the Transporter in a way that may not be a useful factorization for user-defined transport to succeed.

LightStep will consider using either the approximate or the
efficient algorithm for its Span Recorder components, depending on
our experience at building a v3 Golang client library in June 2018.

#### Report Builder component

The Report Builder component encapsulates the encoding logic that
turns concrete interface objects into serialized bytes for
sending.  The Report Builder component typically determines the
concrete types of the data-transport objects that are used.  For
example, if the underlying transport uses protobuf, the Report
Builder is a thin wrapper around a protobuf object.

When user-defined transport is selected, user code becomes
involved in transport and, therefore, there are certain auxiliary
APIs made available for meta-reporting back to LightStep.  The
Report Builder includes an API for setting a pending count of
dropped spans to be sent in the next report.

#### Transporter component

The transporter component handles sending data to LightStep over
an HTTP connection.  A Transporter is only selected in
conjunction with a Span Recorder, when the user has not selected
user-defined transport, so we treat the Transporter as
essentially a "dumb" component.  It should not retry, for
example, that logic must be performed in the Span Recorder, e.g.,
because we have an overall limit on data transmission.

## Span context carrier

LightStep supports several ways to transport ("carry") span context
between applications.  There are several different terms used to
describe this support in OpenTracing, which has led to a confusing,
incomplete matrix of carrier support across client libraries.

- *Carrier format*: describes the logical type of the data, interpreted by the vendor
- *Value type*: describes the type of value passed to and from inject / extract
- *Encoding type*: describes how the context will be encoded, whether as a single base64-encoded header or multiple text headers, for example.

    TODO: Clarify the current state of the world.

    TODO: Specify which combinations of the above must be supported
    for third-generation client libraries.

    TODO: Widen the SpanContext trace_id to 128 bits to match the W3C
    trace propagation spec.
