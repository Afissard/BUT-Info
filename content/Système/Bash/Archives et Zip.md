# tar
Archive un fichier / dossier
## Exemple (tldr)
```
 - [c]reate an archive and write it to a [f]ile:
   tar cf {{path/to/target.tar}} {{path/to/file1 path/to/file2 ...}}

 - [c]reate a g[z]ipped archive and write it to a [f]ile:
   tar czf {{path/to/target.tar.gz}} {{path/to/file1 path/to/file2 ...}}

 - [c]reate a g[z]ipped archive from a directory using relative paths:
   tar czf {{path/to/target.tar.gz}} --directory={{path/to/directory}} .

 - E[x]tract a (compressed) archive [f]ile into the current directory [v]erbosely:
   tar xvf {{path/to/source.tar[.gz|.bz2|.xz]}}

 - E[x]tract a (compressed) archive [f]ile into the target directory:
   tar xf {{path/to/source.tar[.gz|.bz2|.xz]}} --directory={{path/to/directory}}

 - [c]reate a compressed archive and write it to a [f]ile, using the file extension to [a]utomatically determine the compression program:
   tar caf {{path/to/target.tar.xz}} {{path/to/file1 path/to/file2 ...}}

 - Lis[t] the contents of a tar [f]ile [v]erbosely:
   tar tvf {{path/to/source.tar}}

 - E[x]tract files matching a pattern from an archive [f]ile:
   tar xf {{path/to/source.tar}} --wildcards "{{*.html}}"
```
# gzip
Zip un fichier ou un dossier
## Exemple (tldr)
```
 - Compress a file, replacing it with a gzipped compressed version:
   gzip {{file.ext}}

 - Decompress a file, replacing it with the original uncompressed version:
   gzip -d {{file.ext}}.gz

 - Compress a file, keeping the original file:
   gzip --keep {{file.ext}}

 - Compress a file specifying the output filename:
   gzip -c {{file.ext}} > {{compressed_file.ext.gz}}

 - Decompress a gzipped file specifying the output filename:
   gzip -c -d {{file.ext}}.gz > {{uncompressed_file.ext}}

 - Specify the compression level. 1=Fastest (Worst), 9=Slowest (Best), Default level is 6:
   gzip -9 -c {{file.ext}} > {{compressed_file.ext.gz}}
```