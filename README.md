# GoBlaze

GoBlaze ist ein Go-basiertes User-Interface-Framework, das in der Tradition von ASP.NET Razor/Blazor Pages steht – aber als native, serverseitige Lösung ohne Java, Node oder Browser-Build-Toolchain.

## Was ist GoBlaze?

- Eine komponentenbasierte UI-Architektur in Go.
- Serverseitiges Rendering mit VDOM-Diffing und WebSocket-Update-Pipeline.
- Kein eigener JavaScript-Frontend-Build, kein npm, kein Maven, kein separater Backend-Service.
- Deployment als einzelne Go-Binary: einfach, wartbar und leicht zu skalieren.

## Warum GoBlaze?

### Ähnlich wie Razor/Blazor, aber in Go
- Das Konzept erinnert an ASP.NET Razor/Blazor Pages: UI-Logik und Rendering leben eng zusammen.
- UI-Komponenten definieren Struktur, Zustand und Events.
- Der Server rendert die Oberfläche und aktualisiert nur die geänderten Teile.

### Anders als klassische Frontend/Backend-Trennung
- Du brauchst keine „Frontend-App“ in React/Vue und getrennte API-Schicht.
- Die Anwendung bleibt in einer Codebasis: UI, Zustand und Event-Handling sitzen im selben Prozess.
- Das ist besonders sinnvoll für Admin-Dashboards, interne Tools oder Apps, bei denen schnelle Iteration und niedriger Infrastrukturaufwand wichtiger sind als volle SPA-Frontend-Komplexität.

### Single Binary statt Java/Node-Gewitter
- Go erzeugt eine einzelne ausführbare Datei.
- Kein lästiges Java-Hosting, kein npm install, kein Webpack/Vite.
- Updates sind einfache Deployments der neuen Binärdatei.
- Das reduziert Betriebskosten, Komplexität und Abhängigkeiten.

### Skaliert trotzdem
- Go ist für hohe Parallelität und viele Verbindungen gebaut.
- WebSocket-basierte Sessions ermöglichen reaktive UI-Updates ohne komplette Seiten-Reloads.
- Durch horizontale Skalierung hinter einem Load Balancer kann die Plattform wachsen, ohne die Architektur zu sprengen.
- Die serverseitige Zustandsführung erlaubt kleine, effiziente Nachrichten statt riesiger JSON-APIs.

## Für wen ist das geeignet?

- Projekte, die mehr UI-Logik als rohe REST-/GraphQL-APIs benötigen.
- Anwendungen, bei denen die Service- und UI-Logik zusammengehören.
- Entwickler, die native Go-Ausführung und einfache Deployment-Pipelines schätzen.
- Teams, die eine schlanke Alternative zu ASP.NET Blazor Server/WebAssembly oder klassischen JavaScript-Stacks suchen.

## Wichtige Eckpunkte

- Lizenz: `AGPL-3.0-or-later`
- Branch-Schutz: `main` ist als geschützter Branch vorgesehen.
- Alle Quellcode-Dateien enthalten einen AGPL-Hinweis.
- Die vollständige Lizenz steht in `LICENSE`.
- Zusätzliche Hinweise stehen in `AGPL_NOTICE.md`.
- GitHub-Einrichtungsinformationen stehen in `GITHUB-SETUP.md`.

## Branch-Schutz

Um `main` zu sperren, verwende auf GitHub die Branch-Schutz-Regeln:

1. Repository auf GitHub erstellen.
2. Gehe zu `Settings` → `Branches`.
3. Füge eine neue Regel für `main` hinzu.
4. Aktiviere mindestens:
   - `Require a pull request before merging`
   - `Require approvals`
   - `Require status checks to pass before merging`
   - `Restrict who can push to matching branches`

## Lizenz und Kopierschutz

Jedes Veröffentlichen, Verteilen oder Nutzen dieses Projekts setzt voraus, dass die AGPLv3 eingehalten wird.
