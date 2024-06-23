module.exports = {
    extends: [
        "@commitlint/config-conventional"
    ],
    rules: {
        // find meaning and explanations here https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#type
        "type-enum": [2, "always", ["build", "chore", "ci", "docs", "feat", "fix", "refactor", "revert", "style", "test"]],
        "scope-enum": [2, "always", ["chart", "test", "taskfile"]]
    }
}