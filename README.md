# Factorial Code Challenge

For this project, I've used Golang for the Backend API. For the Frontend I've used Golang, HTMX, Tailwind and ChartsJS. And for the database MongoDB. And deployed everything in Docker.

I've used Go throughout this application because it's a language that brings low verbosity and fast results to achieve these goals and it's a language that I'm very familiar with.

I've also included unit tests in both projects for the main `domain` methods

## Frontend
I've decided to use Golang templating feature together with HTMX, which I've imposed on myself this extra work, since I've never used HTMX and heard good things about it paired with Go. Also, I've used React for some time in the past, so I already had this knowledge.

I found it pretty easy to use, just returning those HTML templates already filled with the information that I requested from the backend. Also just using HTML attributes to make GET and POST requests is really easy.

Definetely there was a learning curve there, on how to properly trigger the requests and replacing the target HTML elements correctly. I've had special trouble replacing the tags to seed information to the Javscript that created the Chart. After a lot of back and forths it snapped and just worked.

Also used Tailwind for the styling of the page, since it's pretty effortless to create a nice visual for the page. Just found some examples on the net, adjusted a few things, and it was already done.

For the project, I've separated for the code in an Api module, where I created a generic struct that makes requests to the baseurl and resource passed as parameter. So assuming this could grow in the future, I wouldn't have to create again all the code necessary to make requests to another api. There's also the domain module, where the models to the frontend and a few data manipulation takes place. There's also the log module that it's 2 simple funcs to make it easier to log the application to the stdout. I could've done just the `log.Println` vanilla Go func. There's the resources folder where the HTMLs are stored for the templating.

I've also added different `.env` files for local deploying and docker deploying. Also created a `Dockerfile` to create the image based on the application. `go.mod` and `go.sum` are configuration files for the Go application and `main.go` is where the application is started and the endpoints are set.

## Backend
For the backend it's mainly Go code separated between a few different modules. The database where I made it simple if hypothetically I wanted to query another collection in MongoDB. The domain where I got the data an manipulated it. And the log module that is the same from the frontend.

Besides that it is pretty close the frontend. The `.env` files for environment variables, `Dockerfile` to create the docker image, `go.mod` and `go.sum`, and a `main.go`.

The difference in the `main.go` here is that instead of exposing the HTML templates, I'm returning Json, since it's a REST api.

## Database
For the Database I've used MongoDB, since this is a metrics application, I assume that this would result in millions of unstructured data being inserted, low use case for an update. For this use case a NoSQL database is better.