# Récupération de l'IP et port de Nexus
ARG REPOSITORY

# Image de base
FROM ${REPOSITORY}/javascript:cert

# Ajout de l'app dans le container
WORKDIR /app
COPY ./app /app

# Installation des dépendances
RUN npm i

# Installation du launcher Angular
RUN npm i -g @angular/cli

# Lancement de l'application
CMD ng serve --host 0.0.0.0 --poll 1000 --disable-host-check