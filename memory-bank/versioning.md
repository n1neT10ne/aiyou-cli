# Versionnement et Processus de Release

## Gestion des Versions

Le projet aiyou-cli suit le système de versionnement sémantique (SemVer) avec le format `vX.Y.Z` où :

- **X** : Version majeure - Incrémentée lors de changements incompatibles avec les versions précédentes
- **Y** : Version mineure - Incrémentée lors de l'ajout de fonctionnalités rétrocompatibles
- **Z** : Version de correctif - Incrémentée lors de corrections de bugs rétrocompatibles

### Quand incrémenter chaque partie de la version

#### Version Majeure (X)
Incrémentez la version majeure lorsque vous effectuez des changements incompatibles avec les versions précédentes, tels que :
- Modification de l'interface de ligne de commande qui brise la compatibilité avec les scripts existants
- Suppression de fonctionnalités ou de flags
- Changement du comportement par défaut de manière significative
- Refonte majeure de l'architecture

Exemple : `v1.0.0` → `v2.0.0`

#### Version Mineure (Y)
Incrémentez la version mineure lorsque vous ajoutez des fonctionnalités rétrocompatibles, telles que :
- Ajout de nouveaux flags ou options
- Ajout de nouvelles fonctionnalités
- Améliorations significatives des fonctionnalités existantes
- Changements qui n'affectent pas les utilisations existantes

Exemple : `v1.0.0` → `v1.1.0`

#### Version de Correctif (Z)
Incrémentez la version de correctif lorsque vous effectuez des corrections de bugs rétrocompatibles, telles que :
- Correction de bugs
- Améliorations de performances mineures
- Changements de documentation
- Refactoring interne sans impact sur l'interface utilisateur

Exemple : `v1.1.0` → `v1.1.1`

## Processus de Release

Pour créer une nouvelle release d'aiyou-cli, suivez ces étapes :

1. **Mettre à jour la version dans le code**
   ```go
   // Dans main.go
   const version = "vX.Y.Z"  // Remplacez X.Y.Z par la nouvelle version
   ```

2. **Commiter et pousser les changements**
   ```bash
   git add main.go
   git commit -m "Mise à jour de la version vers vX.Y.Z"
   git push
   ```

3. **Créer et pousser un tag**
   ```bash
   git tag vX.Y.Z
   git push origin vX.Y.Z
   ```

4. **Vérifier le déclenchement du workflow GitHub Actions**
   - Le push du tag déclenchera automatiquement le workflow GitHub Actions
   - Ce workflow va :
     - Construire des binaires natifs pour Windows (.exe) et Linux
     - Créer une release GitHub avec ces binaires
     - Générer des notes de release automatiquement

5. **Vérifier la nouvelle release**
   - Vérifiez que la nouvelle release apparaît sur la [page des releases GitHub](https://github.com/n1neT10ne/aiyou-cli/releases)
   - Vérifiez que les binaires ont été correctement générés et attachés à la release

## Exemple Concret

Pour la version actuelle (v1.1.0), nous avons incrémenté la version mineure car nous avons ajouté deux nouvelles fonctionnalités rétrocompatibles :
1. Support des prompts système multi-lignes dans le fichier de configuration YAML
2. Support de l'entrée standard pour les messages avec le flag `-i/--stdin`

Ces ajouts ne brisent pas la compatibilité avec les utilisations existantes, mais ajoutent de nouvelles fonctionnalités, ce qui justifie une incrémentation de la version mineure.

## Bonnes Pratiques

- Assurez-vous que la version dans le code (`main.go`) correspond toujours à la dernière version taguée
- Incluez des notes de release détaillées pour chaque version
- Testez toujours les changements avant de créer une nouvelle release
- Documentez les changements importants dans le README.md
- Mettez à jour la documentation lorsque vous ajoutez de nouvelles fonctionnalités
