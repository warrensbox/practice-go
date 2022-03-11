

## Part 1:
Given a files of files that needs to be build, return the first set of files that needs to be build prior to building the original file

            A
           / \
          B   C   G
         / \ / \ / 
        D   E   F


onBuild("A") 
Should return D,E,F - to build A, you must first start with D,E and F

## Part 2:
Given a file, what is the next file that can be build.
### Case 1
nextBuild("D")
Should return [] //because you have can't build anything else

nextBuild("E")
Should return [B] //because you have already build "D", so now you can build "B"

nextBuild("F")
Should return [C] //because you have already build "E", so now you can build "C"
### Case 2
nextBuild("D")
Should return [] //because you have can't build anything else

nextBuild("F")
Should return [] //because you have can't build anything else

nextBuild("E")
Should return [B C] //because you have already build "D" and "F", so now you can build "B" and "C"
