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

The main report payload consists of a list of spans.  Span match
the OpenTracing Span specification closely.  Users may apply tag
values to spans downstream in their user-defined transport by
adding them to `lightstep.Span.tags`.

Span tags override reporter tags, making it possible to create spans
with a client on behalf of another process.  To set the reporter GUID
for a span, use:

Key | Meaning
----|--------
`lightstep.guid` | Equivalent to the reporter uuid.  If 64-bits, use a 16-byte hex representation.  If 128-bits, use a uuid.v4 string representation.

## Data Transport 

LightStep client libraries will be factored to include a default
http/2 implementation and a user-defined transport mode.  Although
gRPC is eliminated from third generation client libraries, we use the
same protocol definition and are committed to gRPC support on the
server, meaning that user-defined transport implementations may use
the gRPC [`lightstep.collector.CollectorService.Report`](https://github.com/lightstep/lightstep-tracer-common/blob/4c649d1a7ac52b9cafc7f8d21fe304f8fa4a4ae3/collector.proto#L105) endpoint.

Client libraries may be viewed as an assembly of two parts, one an
implementation of "pure" tracing, and the other a facility to send
data to LightStep.  We will refer to these as the top half and the
bottom half, respectively.

There are several competing interests present when designing client
libraries, which we prioritize as follows:

1. Do not harm the application
1. Do not harm the LightStep service
1. Send as much data as possible
1. Do not waste resources

First and second generation LightStep client libraries offered several
resource- and performance-related parameters, notably: (1) limit on
number of buffered spans, (2) minimum flush period, (3) maximum flush
period, (4) report timeout.  Those libraries were restricted to a
maximum of one simultaneous Report being sent at a time, making these
parameters easy to interpret but difficult to tune.  Third generation
client libraries will be configured by a new set of parameters:

Name               | Interpretation
------------------ | --------------
`max_memory`       | Number of bytes of memory buffered (top)
`max_report_size`  | Limits the size of an outgoing report (bottom)
`max_concurrency`  | Number of CPUs dedicated to sending reports (bottom)
`max_flush_period` | Prevent sending more frequently (bottom)

The `max_memory` parameter applies only to the top half of the
library.  Users should assume that each concurrent sender may use up
to `max_report_size` of memory while sending reports.  The priorities
listed above should be used to determine client behavior, without need
for additional tuning parameters.

#### Note about mobile platforms

We expect that mobile platforms will require separate transport
implementations compared with the same language on a non-mobile
platform.

#### Note about built-in safety

Client libraries are expected to use built-in facilities such as
string formatting and JSON marshalling when producing reports, and are
therefore only as safe as those facilities.  This risk is passed on to
the programmer.

#### Note about protocol buffers vs. JSON

Protocol buffer library support varies significantly by language, and
in some languages there is more than one viable choice of library.
The top half of the library should be not constrain the library used
for encoding span data in the bottom half.

### User-defined transport

The goal of user-defined transport is to insert an abstraction between
the top-half and bottom-half of a client library, allowing the user to
supply a custom bottom half.

This type exists in some of the libraries already, for example a
`Recorder` in
[C++](https://github.com/lightstep/lightstep-tracer-cpp/blob/4ea8bda9aed08ad45d6db2a030a1464e8d9b783f/src/recorder.h#L9)
and a `SpanRecorder` in
[Golang](https://github.com/lightstep/lightstep-tracer-go/blob/644c3d5ecbd0499c50a1329f89ba287921fc1144/options.go#L66),
but the interface is not currently consistent.  Java has multiple transport
options, but no facility for a user-provided transport,
while Objective-C has only a single transport option.

#### User-defined transport: `TracerImpl`

The `TracerImpl` type contains the top-half of the client library,
with access to the user-supplied tracer options (the reporter tags,
access token) and the reporter identity (uuid).

#### User-defined transport: `Recorder`

The `Recorder` interface.

    TODO: Supports `RecordSpan()` and `Flush()`.

#### User-defined transport: `ReportBuilder`

The `ReportBuilder` interface, similar to the existing [C++](https://github.com/lightstep/lightstep-tracer-cpp/blob/4ea8bda9aed08ad45d6db2a030a1464e8d9b783f/src/report_builder.h#L12) interface.

    TODO: Supports `AddSpan()`, `SetPendingClientDroppedSpans()`, `Pending()`.

### Default transport implementation

    TODO: Document and give pseudo-code for managing the tension
    between report size, concurrency level, timeout, and backoff.

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
