---
title: "hbm policy ls"
description: "The policy ls command description and usage"
keywords: [ "policy", "ls" ]
date: "2017-01-27"
menu:
  docs:
    parent: "hbm_cli_policy"
github_edit: "https://github.com/kassisol/hbm/edit/master/docs/reference/commandline/policy_ls.md"
---

```markdown
List policies

Usage:
  hbm policy ls [flags]

Aliases:
  ls, list

Flags:
  -f, --filter value   Filter output based on conditions provided
```

## Filtering

The filtering flag (`-f` or `--filter`) format is a `key=value` pair. If there is more
than one filter, then pass multiple flags (e.g. `--filter "foo=bar" --filter "bif=baz"`)

The currently supported filters are:

* name (policy's name)
* user (username)
* group (group's name)
* resource-type (action|cap|config|device|dns|image|logdriver|logopt|port|registry|volume)
* resource-value (resource's value)
* collection (collection's name)

## Examples

```bash
# hbm policy ls
NAME                GROUP               COLLECTION
policy1             group1              collection1
policy2             group2              collection2
```

```bash
# hbm policy ls -f "user=user1"
NAME                GROUP               COLLECTION
policy1             group1              collection1
```
## Related information

* [policy_add](policy_add.md)
* [policy_find](policy_find.md)
* [policy_rm](policy_rm.md)
