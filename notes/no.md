# TODO

# Done

week one.
completed:

-   cache fix for role based access control
-   improved rbac implementation. able to reduce two db calls to one cache call for the whole rbac process
-   bug fix reported by pavan in usage engine for the case of no data in the db
    {
    "operatorId": "2TQhpDl9JGsh10Ep7Sj3vVUsO7a",
    "userId": "NT",
    "username": "NT",
    "password": "Test@124",
    "role": "telcoAdmin"
    }

-   memory usage
-   cpu
-   http connections
-   gc
-   database profiling
-   balance verification:db

Complete NRF payload that will be supported in NCHF during registration. This documentation is required, as we will be integrating with many different NRF providers

[x] How do we configure the NRF end points in our NCHF interface?

How do we support multiple FQDNs for failover?

What is the heartbeat logic followed with NRF? I remember we will also receive the heartbeat request from NRF and also we will send. So, both the flows should be verified.

What is heart beat interval? How is it configured?

[x] What is the 3GPP release of NRF spec we will support?

1. Poll watching
