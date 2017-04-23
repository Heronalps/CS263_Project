# GO_AWS_Lambda
Spring 2017 CS263 Project repo.

* Goal : Implement profiling for Go AWS Lambda functions.

* Methodology :
> Test, stress, use, and extend this FaaS system (single node and multinode).
>
> Write a set of different go functions from scratch that you use to do something interesting in the system.
>
> Run a series of performance analyses on the example functions and your functions, present the data (resource use, throughput, RTT, etc).
>
> Extend to provide rabbitMQ/celery background tasking service to spawn asynchronous, long running tasks. Determine how to error check, kill, prevent task bombs, and enable other robustness features.

* References:

http://blog.alexellis.io/functions-as-a-service/

https://github.com/alexellis/faas/

https://medium.com/@mweagle/a-go-framework-for-aws-lambda-ab14f0c42cb#.r6qaulxan

https://github.com/eawsy/aws-lambda-go

https://github.com/apex/apex

http://www.avitzurel.com/blog/2016/06/17/writing-an-aws-lambda-function-with-golang/

* Docker image:

> docker run -p 4000:80 heronalps/first_repo:version1.0

> localhost:4000

[[https://github.com/heronalps/GO_AWS_Lambda/docker/service_stack.png|alt=octocat]]
