# LightStep Tracer Client v3 Specification

LightStep Engineering
May 2, 2018

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
third generation libraries remove both of these in favor of raw http/2
transport using either protobuf or JSON encoding.  Third generation
client libraries are expected to perform efficiently in both
high-throughput and low-resource environments.  Third generation
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
should use the available resources efficiently, particularly at the
performance boundary, where we may be forced to drop data or increase
resource usage, depending on service conditions. We identify the most
important scenarios for consideration as follows:

* New spans produced while the buffer of pending data is full
* Non-retryable failures received before the timeout
* Reporting timeout received
* When user code creates excessively large spans or logs

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
interpretation.

We document the interpretation of certain fields here to support
deeper-integration between the application and the mode of
user-defined transport.  To that end, we only document those
fields necessary to support approved uses.  

    TODO: Add documentation comments in the `collector.proto`
    source about the fields that are valid for end-user
    interpretation and modification during user-defined transport.

    TODO: Add documentation comments in the `collector.proto`
    source defining how those fields should be used for the
    implementor of a LightStep client library.

### Report structure

The `lightstep.ReportRequest` message consists of three
user-interpretable fields:

1. **reporter**: describes the client process that generated the batch of spans
2. **auth**:     describes the access token associated with the batch of spans
3. **spans**:    the main payload, a list of span structs

Reports must be smaller than 16MB when serialized as a
binary-encoded protocol message.

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
`lightstep.tracer_platform`         | A string describing the particular client library and transport implementation.
`lightstep.tracer_platform_version` | A string describing the version of the client library.

Reporter fields are effectively applied to each of the spans
contained within the report.  Spans may override each of the
above values by setting the same tag in their own
`lightstep.Span.tags` set.

    TODO: Widen the LightStep reporter guid field to 128 bits.

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
with a client on behalf of another process.  To set the reporter GUID:

Key | Meaning
----|--------
`lightstep.guid` | Equivalent to the reporter uuid, as a 64-bit unsigned decimal string.

    TODO: Specify to represent 128-bit uuid in uuid.v4 format as a string.

## Data Transport 

LightStep client libraries will be factored to include a default
http/2 implementation and a user-defined transport mode.  Although
gRPC is eliminated from third generation client libraries, we use the
same protocol definition and are committed to gRPC support on the
server, meaning that user-defined transport implementations may use
the gRPC [`lightstep.collector.CollectorService.Report`](https://github.com/lightstep/lightstep-tracer-common/blob/4c649d1a7ac52b9cafc7f8d21fe304f8fa4a4ae3/collector.proto#L105) endpoint.

### User-defined transport

    TODO: Add detail section.  Note that the LightStep C++,
    Golang, and Java tracers are already factored for
    user-defined transport.  This exercise is only needed for
    Objective-C.

Note: LightStep will continue to support a gRPC endpoint for
receiving reports.  User-defined transport may continue to use
gRPC for forwarding reports to LightStep.

## Span context carrier

LightStep supports several ways to transport ("carry") span
context between applications.  There are several different terms
used to describe this support in OpenTracing, which have
unfortunately

- *Carrier format*: describes the logical type of the data, interpreted by the vendor
- *Value type*: describes the type of value passed to and from inject / extract
- *Encoding type*: describes how the context will be encoded, whether as a single base64-encoded header or multiple text headers, for example.

    TODO: Specify which combinations of the above must be supported
    for third-generation client libraries.
