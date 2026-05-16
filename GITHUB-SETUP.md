# GitHub Setup / Branch Protection

Dieses Dokument beschreibt, wie du das Repository auf GitHub einrichtest und den Branch `main` schützt.

## Schritt 1: Repository erstellen

1. Erstelle ein neues Repository auf GitHub unter deinem Account.
2. Wähle den Namen `goblaze` oder einen anderen passenden Namen.
3. Füge keine zusätzliche Lizenz hinzu, da die Lizenz bereits in dieser Codebasis enthalten ist.

## Schritt 2: Remote einrichten

Führe lokal im Projektordner aus:

```bash
git remote add origin git@github.com:<dein-account>/<repo-name>.git
git push -u origin main
```

## Schritt 3: Branch-Schutz aktivieren

In GitHub:

1. Gehe zu `Settings` → `Branches`.
2. Klicke auf `Add rule`.
3. Gib `main` als Branch-Namen ein.
4. Aktiviere:
   - `Require a pull request before merging`
   - `Require approvals`
   - `Require status checks to pass before merging`
   - `Restrict who can push to matching branches`
5. Speichere die Regel.

## Optional: GitHub CLI

Falls du die GitHub CLI benutzt, kann die Branch-Protect-Regel wie folgt erstellt werden:

```bash
# GitHub CLI muss installiert und angemeldet sein
gh repo create <dein-account>/<repo-name> --public --source=. --remote=origin --push
```

Branch-Schutz muss derzeit manuell in den Einstellungen eingerichtet werden.
