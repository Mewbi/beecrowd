animal = ""
c1 = input()
c2 = input()
c3 = input()

animals = {
    "vertebrado": {
        "ave": {
            "carnivoro": "aguia",
            "onivoro": "pomba",
        },
        "mamifero": {
            "onivoro": "homem",
            "herbivoro": "vaca",
        },
    },
    "invertebrado": {
        "inseto": {
            "hematofago": "pulga",
            "herbivoro": "lagarta",
        },
        "anelideo": {
            "hematofago": "sanguessuga",
            "onivoro": "minhoca",
        },
    },
}

if c1 in animals and c2 in animals[c1] and c3 in animals[c1][c2]:
    animal = animals[c1][c2][c3]

print(animal)
