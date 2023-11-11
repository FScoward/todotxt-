- プロジェクトとはタスクである
- タスクとはゴールである
- 1 段目はプロジェクト名

  - 2 段目以降はタスクである

- TYPE TITLE CDATE UDATE

- TYPE には PRJ,TODO,DONE,DOC,LOG

- PRJ タスク管理プロジェクト 2023-10-31
  - DOC xxxxx 2023-10-31
  - TODO xxxxxx 2023-10-31

## 仕様

1 ファイル 1 プロジェクトとする。

```ProjectX.md
---
projectName: ProjectX
id: xxxxxx
---

## Task

- TODO task1
  - TODO sub task2
    ID: xxx-xxx-xxx
    CREATED: 2023-10-10
  - DONE sub task3
    ID: yyy-yyy-yyy
    CREATED: 2023-10-10
    CLOSED: 2023-11-10

## Doc

## Note

## Meeting

- 2023-10-12 10:00 X Meeting

## Log

- Add xxx-xxx-xxx 2023-10-10
- Add yyy-yyy-yyy 2023-10-10
- Done yyy-yyy-yyy 2023-11-10

```
