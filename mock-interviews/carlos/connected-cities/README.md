Connected Cities

   There are n cities. Some of them are connected, while some are not. If city a
   is connected directly with city b, and city b is connected directly with city
   c, then city a is connected indirectly with city c.

   A province is a group of directly or indirectly connected cities and no other
   cities outside of the group.

   You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the
   ith city and the jth city are directly connected, and isConnected[i][j] = 0
   otherwise.

   Return the total number of provinces.

   The following city configuration which counts as 2 provinces:

   0 <-> 1
   2

   Is represented as follows:

   [
    0 : [1, 1, 0], // 0
    1:  [1, 1, 0], // 1
    2:  [0, 0, 1], // 2
   ]

   The following city configuration which counts as 1 province:

   0 <-> 1 <-> 2

   Is represented as follows:

   [
     [1, 1, 0], // 0
     [1, 1, 1], // 1
     [0, 1, 1], // 2
   ]

   The following city configuration which counts as 3 provinces:

   0
   1
   2

   Is represented as follows:

   [
     [1, 0, 0], // 0
     [0, 1, 0], // 1
     [0, 0, 1], // 2
   ]
