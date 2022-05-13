# go run cmd/main.go -conf cmd/conf.yaml

# crons 任务主体
# crons_log 任务执行记录

# 三个接口
# 任务详情      /cron/info      参数 id 
# 任务列表      /cron/list      参数 page，pageSize
# 任务批量新建  /cron/batch     参数 list(详见test)
# 任务记录列表  /cron_log/list  参数 page，pageSize