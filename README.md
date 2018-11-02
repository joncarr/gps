# Go Package Search (GPS)

Go Package Search (gps) provides a means to search godoc.org via the command-line.  It's extremely simple to use with logical options. Currently, no flags additional flags are available.

## Commands
    - gps (root command)
        - help
        - find
            - limit
            - sort

## Usage
To use, simply call the program and provide a search term for the `find` sub-command. `find` currently has 2 flags where you can pass some options. First, you can limit your search results by passing `-l` or `--limit` and providing a number between 1 and 100.  By default, 100 is used since the api returns 100 results. You can also sort the results alphabetically by name, by score (descending), by imports (descending), and by number of repo stars (descending).  You sort by passing find `-s` or `--sort` and providing one of 4 strings.  The available options are:  

    alpha: to sort alphabetically by package name
    score: to sort by score
    stars: to sort by stars
    imports: to sort by imports (***Default Option***)
  
Search with all defaults - Sorted by imports, 100 results:  
    
    `gps find [search-term]`  

Search and provide sort method - Sorted by stars, 100 results:  
    
    `gps find [search-term] -s stars`
    `gps find [search-term] --sort stars`

Search and provide result limit - Sorted by imports, 8 results:  
    
    `gps find [search-term] -l 8`
    `gps find [search-term] --limit 8`

Search and provide sort method and result limit - Sorted by score, 15 results:  
    
    `gps find [search-term] -s score -l 15`
    `gps find [search-term] --sort score --limit 15`


That's it!
 
  
  ![gps image](https://raw.githubusercontent.com/joncarr/gps/master/extras/gps_screen.png)
  
You will be provided the results of the search in your terminal.  Each search result has an entry number associated with it. Just enter the entry number, hit `enter` and the package will install. You may also enter `q` to quit the program if you do not want to install anything and want to quit the program.

## Future "Maybes"
More to come...
