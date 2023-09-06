# Twitter Clone in Go

Twitter clone develop in microservice architecture written in Go.

Services:

-   Front end
-   Broker
    -   Single entry point of backend.
    -   Control routing to other microservices.
-   Session
-   Authentication
-   Post
    -   Create, update and delete posts.
    -   Interact with others' posts.
-   Social
    -   Manage follow / subscriptions.
