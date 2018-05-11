# Journey Project

L'idée est de partir d'une grande question existentielle : qu'elles sont mes chances de victoire quand je commence une bataille à Risk.
Une fois que j'aurai codé ça, je me pencherai sur l'implémentation d'une map. et puis sur une mécanique de jeu (pas forcément fidèle au Risk). À terme j'aimerait avoir un backend assez bien designé pour pouvoir, d'ici quelques dizaines d'heures de travails, implémenter un frontend.
L'agilité voudrait que j'iter dès le début avec ma partie frontend. Seulement c'est le backend qui m'intéresse dans un premier temps et je n'ai pas envie de perdre du temps sur le frontend en ce début de projet.

#### Combat type Risk

Evolution du pourcentage de chance de gagner contre 100 défenseurs en fonction du nombre d'attaquant.
![N|Solid](https://raw.githubusercontent.com/haagor/Journey_Project/master/img/offensevs100.png)

Evolution du pourcentage de chance de gagner en fonction du nombre attaquant=défenseur.
![N|Solid](https://raw.githubusercontent.com/haagor/Journey_Project/master/img/armytoarmy.png)

#### Map

Je distingue 3 catégories d'implémentation de map :

1. Découpage en zone, accéssible de proche en proche, comme dans Risk.

![N|Solid](https://raw.githubusercontent.com/haagor/Journey_Project/master/img/Risk.png)
  

2. Cadrillage, que ce soit carré ou hoctogonal.

![N|Solid](https://raw.githubusercontent.com/haagor/Journey_Project/master/img/HMM.png)
  

3. Repère orthonormé.

![N|Solid](https://raw.githubusercontent.com/haagor/Journey_Project/master/img/BB.png)
  
Le choix numéro 2 me semble le plus approrié : plus simple à implémenter que la mpa en coordonnée, et proposant un meilleur résultat en terme de navigation que le découpage en zone.


  ***
`todo` implem mvp map

*Simon P.*