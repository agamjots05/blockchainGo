Some things to note
1. When you create an `internal` directory, this can only be imported by the parent project. 
    - So if users wanted to use packages (e.g., `auth`) which is located within said `internal` package. They would be unable to do so.

The main point of using this package is to encapsulate data and dis-allow access from specific sources.

