# Récupération de l'IP et port de Nexus
ARG REPOSITORY

# PORT serveur
ARG PORT
ENV PORT ${PORT}

# Image de base
FROM ${REPOSITORY}/python:cert

# Ajout de l'app dans le container
WORKDIR /app
COPY ./app /app

COPY requirements.txt /app/

RUN pip install --upgrade pip
# Installation de Django
RUN pip install django

# Installation des dépendances
RUN pip install -r requirements.txt

# Commande de base de l'application
CMD python manage.py runserver 0:$PORT