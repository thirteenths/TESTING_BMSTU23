Feature: User Authentication

  Scenario: User authentication
    When I send POST request to "/users/login"
    Then the response code should be 200
    And the response should match json:
      """
     "Code send on email"
      """

    When I send POST request to "/users/verify"
    Then the response code should be 200
    And the response should match json:
      """
     "Code successful"
      """