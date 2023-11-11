package parser

import (
	"strings"
	"todotxtplus/internal/model"
)

func parseMetadata(lines []string) (model.Metadata, error) {
	// メタデータを格納するための変数を初期化
	metadata := make(model.Metadata)
	inMetadataSection := false

	// 各行を解析
	for _, line := range lines {
		// メタデータセクションの開始または終了を検出
		if strings.TrimSpace(line) == "---" {
			if inMetadataSection {
				// 2回目の "---" ならメタデータセクションの終了
				break
			}
			// メタデータセクションの開始をマーク
			inMetadataSection = true
			continue
		}

		// メタデータの解析ロジック
		// 例: "key: value" 形式の行を解析
		if inMetadataSection {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				metadata[key] = value
			}
		}
	}

	// メタデータとnilエラーを返す
	return metadata, nil
}

func parseTasks(lines []string) ([]model.Task, error) {
	var tasks []model.Task
	// 現在処理中のタスクのリスト
	var currentTasks *[]model.Task = &tasks
	var currentTask model.Task

	// タスクの階層構造の管理。全てのタスクを階層構造で持つ。
	taskStack := []*[]model.Task{currentTasks}

	lastIndentLevel := 0

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		// インデントレベルに基づいて現在のタスクリストを調整
		// if lastIndentLevel < len(taskStack) {
		// 	taskStack = taskStack[:lastIndentLevel+1]
		// }
		currentTasks = taskStack[lastIndentLevel]

		// タスクまたはサブタスクの行を解析
		if strings.HasPrefix(trimmedLine, "- TODO") || strings.HasPrefix(trimmedLine, "- DONE") {
			indentLevel := (len(line) - len(trimmedLine))

			status := strings.Fields(trimmedLine)[1]
			description := strings.Join(strings.Fields(trimmedLine)[2:], " ")

			newTask := model.Task{
				Description: description,
				Status:      status,
				SubTasks:    []model.Task{},
				Details:     make(map[string]string),
			}

			// インデントが下がっているのであればSubTaskへの追加
			if lastIndentLevel < indentLevel {
				// parentTask := &tasks[len(tasks)-1]
				parentTask := &(*currentTasks)[lastIndentLevel]
				parentTask.SubTasks = append(parentTask.SubTasks, newTask)
			} else {
				*currentTasks = append(*currentTasks, newTask)
			}

			// TODO 現在位置の修正（現在取り扱っているタスクのルートの再設定）

			// *currentTasks = append(*currentTasks, newTask)

			// タスクスタックを更新
			taskStack[lastIndentLevel] = currentTasks
			// if len(taskStack) > indentLevel+1 {
			// 	taskStack[indentLevel+1] = &(*currentTasks)[len(*currentTasks)-1].SubTasks
			// } else {
			// 	taskStack = append(taskStack, &(*currentTasks)[len(*currentTasks)-1].SubTasks)
			// }

			currentTask = newTask
			// lastIndentLevel = indentLevel
		} else if strings.HasPrefix(line, "\t") {
			// タスクの詳細情報を解析
			parts := strings.SplitN(trimmedLine, ": ", 2)
			if len(parts) == 2 {
				// currentTask := &(*currentTasks)[lastIndentLevel]
				currentTask.Details[parts[0]] = parts[1]
			}
		}

	}

	return tasks, nil
}
