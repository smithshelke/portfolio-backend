## Project Overview

This is a Go project with the module name `shelke.dev/api`. It is likely an API service given the module name. The project uses Go version `1.24.4`.
I am creating the backend for the portfolio-ui located in $HOME. There are several usecases in the portfolio project.
One of the usecase is task management.
It contains models like Feature, Task, User, etc. 
The user can create a feature, add tasks for that feature and other things such as assigned user, priority and status of the task.
The user can also specify what github repos that are linked to the task/feature.
I also want to integrate google gemini apis so that it can automatically read the task, feature description, github repos.
Gemini or any other AI api can read this and write the code and raise a PR in the repos.



## Building and Running

**Building:**

To build the project, navigate to the project root directory and run:

```bash
go build ./...
```

**Running:**

To run the project, you would typically execute the compiled binary. The exact command will depend on the main package location and the desired execution method. A common way to run a Go application is:

```bash
go run .
```

*(TODO: Add more specific instructions for running the application, including any required environment variables or configuration.)*

**Testing:**

To run tests for the project, use the following command:

```bash
go test ./...
```

## Development Conventions

*(TODO: Add information about coding style, testing practices, and contribution guidelines. This would typically involve looking for files like `CONTRIBUTING.md`, `.golangci.yml`, or specific patterns in the codebase.)*
Use the built in http libraries for the server. 
Use the built in packages as much as possible.
Inform me if there is need for an external package
Use the hexagonal pattern for developement

## Database Schemas

**Tasks Table:**
fields: id, name, description, created_at, updated_at, created_by, feature_id, feature_name, priority, status, git_data 

**Feature Table:**
fields: id, name, description, created_at, updated_at, created_by, priority, status

**Feature Owners Table:**
fields: id, user_id, feature_id, user_name, user_role

**User Table:**
fields: id, name, role, created_at, updated_at, created_by