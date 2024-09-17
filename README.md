# S5 Sae Emulateur

## Auteurs

- Grégoire Launay-Bécue
- Gwendal Margely

# RISC-V Emulateur

Ce projet est un émulateur de processeur RISC-V écrit en Go. Il permet de décoder et de désassembler des instructions RISC-V à partir de fichiers binaires.

## Table des matières

- [Introduction](#introduction)
- [Structure du projet](#structure-du-projet)
- [Déroulé de la création des éléments](#déroulé-de-la-création-des-éléments)
- [Décisions prises](#décisions-prises)
- [Problèmes rencontrés](#problèmes-rencontrés)
- [Utilisation](#utilisation)
- [Contributeurs](#contributeurs)

## Introduction

Ce projet a été réalisé dans le cadre de la SAÉ du semestre 5 du BUT Informatique, Parcours Réalisation d’applications : conception, développement, validation. L'objectif est de développer un émulateur de processeur RISC-V en utilisant le langage de programmation Go.

## Structure du projet

``` bash
s5-sae-emulateur/
│ 
├── main.go
├── decode.go
├── disassemble.go
├── utils.go
├── opcodes.json
├── README.md
├── go.mod
├── main.exe
└── ...
```


## Déroulé de la création des éléments

### 1. Création de la structure de base du projet

Nous avons commencé par créer la structure de base du projet avec les fichiers suivants :
- `main.go` : Le point d'entrée principal du programme.
- `decode.go` : Contient les fonctions pour lire le fichier binaire, extraire les mots de 32 bits, et décoder les instructions.
- `disassemble.go` : Contient les fonctions pour désassembler les instructions RISC-V en langage assembleur.
- `utils.go` : Contient des fonctions utilitaires, comme l'extension des valeurs immédiates et la lecture du fichier binaire.
- `opcodes.json` : Contient la table de correspondance des opcodes en format JSON.
- `README.md` : Ce fichier de documentation.

### 2. Lecture et compréhension de la spécification RISC-V

Nous avons lu et compris les sections pertinentes de la spécification RISC-V, notamment les jeux d'instructions RV32I. Nous avons également compris les types d'encodage (R, I, S, U, UJ, SB).

### 3. Développement de la fonction d'extraction d'opcode

Nous avons implémenté une fonction en Go pour extraire les 7 bits d'opcode d'un mot de 32 bits. Cette fonction est située dans `utils.go`.

### 4. Création de la table de correspondance des opcodes

Nous avons créé un fichier JSON (`opcodes.json`) pour stocker la table de correspondance des opcodes. Ce fichier contient les opcodes et leurs types d'encodage correspondants.

### 5. Développement de la fonction de décodage des instructions

Nous avons implémenté une fonction pour lire le fichier JSON de la table des opcodes et utiliser cette table pour déterminer le type d'encodage des instructions. Cette fonction est située dans `decode.go`.

### 6. Développement de la fonction de désassemblage des instructions

Nous avons implémenté des fonctions pour désassembler les instructions RISC-V en langage assembleur. Ces fonctions sont situées dans `disassemble.go`.

### 7. Intégration et tests

Nous avons intégré toutes les fonctions développées dans le projet Go. Nous avons testé le programme avec des fichiers binaires d'exemple pour s'assurer que le décodage et le désassemblage fonctionnent correctement.

## Décisions prises

### 1. Utilisation de Go

Nous avons choisi d'utiliser Go pour ce projet en raison de sa performance et de sa simplicité. Go est un langage compilé vers du code natif, ce qui est idéal pour un émulateur de processeur.

### 2. Utilisation d'un fichier JSON pour la table des opcodes

Nous avons décidé d'utiliser un fichier JSON pour stocker la table de correspondance des opcodes. Cela rend le code plus modulaire et facile à maintenir, car nous pouvons facilement ajouter ou modifier des opcodes sans toucher au code source.

### 3. Structure modulaire

Nous avons structuré le projet de manière modulaire en séparant les différentes parties du code dans des fichiers distincts. Cela rend le code plus lisible et plus facile à maintenir.

## Conclusion

Nous avons réussi à implémenter les fonctions de base pour décoder et désassembler les instructions RISC-V. Le projet est maintenant prêt pour des tests plus approfondis et des améliorations futures.

