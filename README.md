WP User Enum - WordPress User Enumeration
=========================================

End goal is to be able to enumerate WordPress users on a site, using different techniques:

1. JSON API - direct
2. JSON API - circumvented
3. Old-school URL structure enumeration


Usage
-----

```
Usage:
  -url string
        WordPress URL. Required.
  -enum int
        Enumeration type. One of 0, 1 or 2:
		0: JSON API, normal REST endpoint
		1: JSON API, via GET query var
		2: brute guessing via author param
  -start int
        Start enumeration at this user ID (default 1). Used with enum type 2.
  -end int
        End enumeration with this user ID (default 10). Used with enum type 2.
  -pretty
        Pretty-print the results
  -ua
        Randomize User-Agent
```
