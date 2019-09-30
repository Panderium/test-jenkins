# Installation des dépendances pour contruire l'application
### Selon les versions de python
```
sudo apt install python-pip
sudo apt install python3-pip
```
```
sudo pip install django
sudo pip3 install django
```

# Commandes du starter Django

## Ne pas oublier de changer le port d'écoute dans le .env

### Création d'une application Django
```
./create-django-app.sh
```

### Lancement du container en forcent le build
```
docker-compose up -d --build
```

### Lancement du container
```
docker-compose up -d
```

### Lancer une commande dans le container via un terminal (Linux)
```
docker exec -i [container_name] [command]
```
### Exemple
```
docker exec -i mon_container pip -r requirements.txt
```

# Mise en preprod et prod
Afin de mettre l'application en preprod et prod il est recommandé d'utiliser Gunicorn et Nginx.
L'application doit être configurée pour la production (ajout de fichier de conf pour Gunicorn et Nginx, suppression du Debug...) et le fichier Dockerfile.production rempli en conséquence.