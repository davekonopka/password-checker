# 20 Add Taskfile

## Prompt

Could you help me create a go-task Taskfile.yml for this project?

## Response

Sure! go-task is a task runner / simpler Make alternative written in Go. Here's how you might set up a Taskfile.yml for your project:

```yaml
version: '3'

tasks:
  build:
    desc: Build the project
    cmds:
      - go build -o main .

  test:
    desc: Test the project
    cmds:
      - go test -v ./...

  docker:
    desc: Build Docker image
    cmds:
      - docker build -t password-checker .

  run:
    desc: Run the project
    cmds:
      - go run .

  docker-compose:
    desc: Start services with docker-compose
    cmds:
      - docker-compose up -d

  down:
    desc: Stop services with docker-compose
    cmds:
      - docker-compose down

  all:
    desc: Build, test, and run the project
    cmds:
      - task: build
      - task: test
      - task: run
```

You can run these tasks with the command task <task-name>. For example, to build the project, you would run task build.

Some tasks like all are composed of other tasks, so running task all would execute the build, test, and run tasks.

Please adjust the Taskfile as per your needs.
