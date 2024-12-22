# find
Cherche et trouve les dossier / fichier dans le dossier donné
## Exemples (tldr)
```
 - Find files by extension:
   find {{root_path}} -name '{{*.ext}}'

 - Find files matching multiple path/name patterns:
   find {{root_path}} -path '{{**/path/**/*.ext}}' -or -name '{{*pattern*}}'

 - Find directories matching a given name, in case-insensitive mode:
   find {{root_path}} -type d -iname '{{*lib*}}'

 - Find files matching a given pattern, excluding specific paths:
   find {{root_path}} -name '{{*.py}}' -not -path '{{*/site-packages/*}}'

 - Find files matching a given size range, limiting the recursive depth to "1":
   find {{root_path}} -maxdepth 1 -size {{+500k}} -size {{-10M}}

 - Run a command for each file (use {} within the command to access the filename):
   find {{root_path}} -name '{{*.ext}}' -exec {{wc -l {} }}\;

 - Find files modified in the last 7 days:
   find {{root_path}} -daystart -mtime -{{7}}

 - Find empty (0 byte) files and delete them:
   find {{root_path}} -type {{f}} -empty -delete
```
# grep
Recherche une chaine de caractère dans un fichier
## Exemples (tldr)
```
 - Search for a pattern within a file:
   grep "{{search_pattern}}" {{path/to/file}}

 - Search for an exact string (disables regular expressions):
   grep --fixed-strings "{{exact_string}}" {{path/to/file}}

 - Search for a pattern in all files recursively in a directory, showing line numbers of matches, ignoring binary files:
   grep --recursive --line-number --binary-files={{without-match}} "{{search_pattern}}" {{path/to/directory}}

 - Use extended regular expressions (supports "?", "+", "{}", "()" and "|"), in case-insensitive mode:
   grep --extended-regexp --ignore-case "{{search_pattern}}" {{path/to/file}}

 - Print 3 lines of context around, before, or after each match:
   grep --{{context|before-context|after-context}}={{3}} "{{search_pattern}}" {{path/to/file}}

 - Print file name and line number for each match with color output:
   grep --with-filename --line-number --color=always "{{search_pattern}}" {{path/to/file}}

 - Search for lines matching a pattern, printing only the matched text:
   grep --only-matching "{{search_pattern}}" {{path/to/file}}

 - Search
   stdin
 for lines that do not match a pattern:
   cat {{path/to/file}} | grep --invert-match "{{search_pattern}}"
```
# test
Compare deux fichiers et leurs contenus
## Exemple (tldr)
```
 - Test if a given variable is equal to a given string:
   test "{{$MY_VAR}}" == "{{/bin/zsh}}"

 - Test if a given variable is empty:
   test -z "{{$GIT_BRANCH}}"

 - Test if a file exists:
   test -f "{{path/to/file_or_directory}}"

 - Test if a directory does not exist:
   test ! -d "{{path/to/directory}}"

 - If A is true, then do B, or C in the case of an error (notice that C may run even if A fails):
   test {{condition}} && {{echo "true"}} || {{echo "false"}}
```
# uniq
Donne les lignes unique d'un fichier donnée (donc les ligne qui ne sont pas répété)
## Exemples (tldr)
```
 - Display each line once:
   sort {{path/to/file}} | uniq

 - Display only unique lines:
   sort {{path/to/file}} | uniq -u

 - Display only duplicate lines:
   sort {{path/to/file}} | uniq -d

 - Display number of occurrences of each line along with that line:
   sort {{path/to/file}} | uniq -c

 - Display number of occurrences of each line, sorted by the most frequent:
   sort {{path/to/file}} | uniq -c | sort -nr
```
# sort
Trie les ligne du fichier donné en paramètre
## Exemple (tldr)
```
 - Sort a file in ascending order:
   sort {{path/to/file}}

 - Sort a file in descending order:
   sort --reverse {{path/to/file}}

 - Sort a file in case-insensitive way:
   sort --ignore-case {{path/to/file}}

 - Sort a file using numeric rather than alphabetic order:
   sort --numeric-sort {{path/to/file}}

 - Sort "/etc/passwd" by the 3rd field of each line numerically, using ":" as a field separator:
   sort --field-separator={{:}} --key={{3n}} {{/etc/passwd}}

 - Sort a file preserving only unique lines:
   sort --unique {{path/to/file}}

 - Sort a file, printing the output to the specified output file (can be used to sort a file in-place):
   sort --output={{path/to/file}} {{path/to/file}}

 - Sort numbers with exponents:
   sort --general-numeric-sort {{path/to/file}}
```
# wc
Compte les lignes / mots / bytes d'un fichier
## Exemple (tldr)
```
 - Count all lines in a file:
   wc --lines {{path/to/file}}

 - Count all words in a file:
   wc --words {{path/to/file}}

 - Count all bytes in a file:
   wc --bytes {{path/to/file}}

 - Count all characters in a file (taking multi-byte characters into account):
   wc --chars {{path/to/file}}

 - Count all lines, words and bytes from
   stdin
:
   {{find .}} | wc

 - Count the length of the longest line in number of characters:
   wc --max-line-length {{path/to/file}}
```