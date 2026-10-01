Feature: Greeting

  @story-1
  Rule: A supplied name is included in the greeting

    Scenario: Calling hello with a name
      Given the greeter service is running
      When an API Caller calls "GET /hello?name=Alice"
      Then the response is a JSON greeting that includes "Alice"

  @story-2
  Rule: A missing or empty name still returns a usable greeting

    Scenario: Calling hello with no name parameter
      Given the greeter service is running
      When an API Caller calls "GET /hello" with no name parameter
      Then the response is a JSON greeting that includes "World"

    Scenario: Calling hello with an empty name parameter
      Given the greeter service is running
      When an API Caller calls "GET /hello?name=" with an empty name
      Then the response is a JSON greeting that includes "World"
