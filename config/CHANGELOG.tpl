# {{#environmentVariable}}HELM_CHART_NAME{{/environmentVariable}}

{{#releases}}
## [{{name}}](https://github.com/curuvija/charts/releases/{{name}}) ({{date}})

{{#sections}}
### {{name}}

{{#commits}}
* [{{#short5}}{{sha}}{{/short5}}](https://github.com/curuvija/charts/commit/{{sha}}) {{message.shortMessage}} ({{authorAction.identity.name}})

{{/commits}}
{{^commits}}
No changes.
{{/commits}}
{{/sections}}
{{^sections}}
No changes.
{{/sections}}
{{/releases}}
{{^releases}}
No releases.
{{/releases}}