# GO apigateway blueprint by Saif Hamdan

## Getting Started

Follow these steps to set up and run the apigateway blueprint project.

### Step 1: Init Project

Run the following command to install the modules and the tools you need to work with it. 

```sh
make init
```

#### Step 2: Make your own .env file

make your own .env file and store your credentials there.

### Step 1: Export Environment Variables

Run the following command to export the necessary environment variables:

```sh
source export.sh
```

### Step 2: Start the Docker Stack

```sh
make stackup
```

### Step 3: Seed the Database

if you didn't seed your database before run this command otherwise skip it.
or it will duplicate the data you already have

```sh
make seed
```

### Step 4: Run the Application

```sh
make run
```
