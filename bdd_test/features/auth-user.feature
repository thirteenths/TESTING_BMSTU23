Feature: User Authentication

  Scenario: User authentication
    When On request "/user/login" I send json:
        """
        {
  		"email": "rachelle.huel@ethereal.email",
  		"password": "C6s2S9qe6WrTMB7z3u"
	    }
        """
    Then the response code should be 200
    And the response should match json:
        """
          Code send on email
        """

    When On request "/user/verify" I send json:
        """
        {
  		"email": "rachelle.huel@ethereal.email",
  		"code": "23456789"
	    }
        """
    Then the response code should be 200
    And the response should match json:
        """
          Code successful
        """

  Scenario: User Update Password
    When On request "/user/update" I send json:
        """
        {
  		"email": "rachelle.huel@ethereal.email",
  		"password": "C6s2S9qe6WrTMB7z3u"
	    }
        """
    Then the response code should be 200
    And the response should match json:
        """
          Code send on email
        """

    When On request "/user/verify" I send json:
        """
        {
  		"email": "rachelle.huel@ethereal.email",
  		"code": "23456789"
	    }
        """
    Then the response code should be 200
    And the response should match json:
        """
          Code successful
        """