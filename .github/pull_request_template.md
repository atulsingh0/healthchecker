## Describe what you have changed and why

*Explain what this PR hopes to achieve.*

*Think about the robustness and reliability principles. Ideally you've already talked through your design with another engineer, if not is that worth doing now?*

## Please do the following before requesting review

*Things the linter / automatic tooling will pick up are not included here*

- [ ] If there are data-store impacting changes, will they be safe during the rollout?
- [ ] Ensure there are added / changed tests at the appropriate level (e2e, integration, unit).
- [ ] Commented any code that necessarily (e.g. for performance reasons) isn't obvious.
- [ ] If adding a new dependency, have you checked the licence, issue history, maintainer attitude to issues, timeliness to review PRs.
- [ ] If adding a new data store query, make sure it's indexed.

## When can this be deployed?

If any of the following are true, please write down your progress / work somewhere linked to this PR (ticket or comments in this PR).

- [ ] Does support comms / external comms need to go out first? (e.g. this changes user-facing behaviour).
- [ ] Does it require canary?
- [ ] Does it depend on another PR? If so list it.
- [ ] Is there are particular chart / dashboard / other system that should be monitored during rollout?
