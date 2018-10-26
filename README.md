# Go Package Search (GPS)

Go Package search (gps) provides a means to search godoc.org via the command-line.  It's extremely simple to use as there is only one command. Currently, no flags additional flags are available.

## Usage
To use, simply call the program and provide a search term.  
  
```gps [search-term]```  
  
That's it!

As an example:
  
```gps twitter```  
  
You will be provided the results of the search in your terminal.  Each search result has an entry number associated with it. Just enter the entry number, hit `enter` and the package will install. You may also enter `q` to quit the program if you do not want to install anything and want to quit the program.

## Future "Maybes"
The result list displays in the same order the API returns the results.  I'd like to add some flags to sort the list by different criteria (score, stars, imports...)
