flower(rose,red).
flower(marigold,yellow).
flower(tulip,red).
flower(daffodil,yellow).
flower(rose,yellow).
flower(maigold,red).
flower(rose,white).
flower(cornflower,purple).
%CASEI)Two Reds
bouquet(L):-flower(X,red),
			flower(Y,red),
			X \==Y,
			flower(Z,_),
			L = [X,Y,Z].
%CASEII)ALL Different
bouquet(L):-!,
			flower(X,C1),
			flower(Y,C2),
			flower(Z,C3),
			C1 \== C2,
			C2 \== C3,
			C3 \== C1,
			L =[X,Y,Z].