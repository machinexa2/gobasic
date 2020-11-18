package gobasic


func InArray(check string, list []string) bool {
        for _, individual := range list {
                if individual == check {
                        return true
                }
        }
        return false
}

