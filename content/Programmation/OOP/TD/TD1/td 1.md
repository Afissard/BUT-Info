# Exercice 1 : compilation et éxécution
1. compiler un .kt vers /bin : `kotlinc src/Main1.kt -o /bin`
2. Compiler d'abord Dep11.kt puis Main1.kt : `kotlinc src/Dep11 -d bin && kotlinc -cp bin src/main1.kt -d bin
3. Compiler deux fichier en même temps : `kotlinc src/Main1.kt -d bin cp bin` ou `kotlinc -cp /Dep11 src/Main1.kt -d bin`
4. exécuté Main1Kt : `kotlin Main1Kt -cp bin`
5. Créer un .jar : `kotlinc src/Dep11Kt scr/Main1.kt -d out/main1.jar include-runtime`
6. exécuté le .jar : `kotlin out/main1.jar` ou `java -jar out/main1.jar`
7. exécuté la classe Main1Kt présente dans main1.jar : `kotlinc cp out/main1.jar Main1Kt`
8. bin est vide, compiler Main2.kt dans bin : `kotlinc -cp class -bin other_src/Main2.kt`
9. Comment exécuter Main2.kt : `kotlin -cp class:bin Main2.kt`
10. bin est vide, compiler Test1.kt dans bin : `kotlinc test/Test1.kt -d bin -cp jar/junit.jar`
11. Compiler Test12.kt dans bin :  `kotlinc -cp jar.junit:out.Main1.kt test/Test12.kt -d bin`
12. bin est vide, compilé Main3.kt et produire un Jar nommé main2.jar dans bin : `kotlinc -cp jar/junit.jar:out/Main1.kt:class -d bin/main2.jar include-runtime` 
13. Comment exécuté la classe Main3Kt présente dans le jar main3.jar : `kotlin -cp jar/junitjar:out/Main1.jar:class:bin/Main2.jar`
# Exercice 2 :  variables et référence et leurs representation dans la mémoire
