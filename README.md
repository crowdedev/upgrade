# Crowde Upgrader

## Motivation

Deployment to production is hard step when we have many changes in database including create new table, remove unused table, add new record, makesurement that records in some table is not duplicate and many others step.

We need to make sure each processes are predictable, maintainable, easy to trace and separate between checking condition and executing upgrade to prevent human errors

## Install

`go get github.com/crowdeco/upgrade`

## Usage

```go
type AddNewRecordInSomethingTable struct {
    DB *gorm.DB
}

func (a *AddNewRecordInSomethingTable) Upgrade() {
    //AddNewRecordInSomethingTable Logic without need to check if record exist or not because checking is handled by Support() function
}

func (a *AddNewRecordInSomethingTable) Support() bool {
    a := checkIfRecordExistInSomethingTable()
    if a {
        return false
    }

    return true
}

func (a *AddNewRecordInSomethingTable) Order() int {
    return 1
}
```

```go
upgrade := AddNewRecordInSomethingTable{DB: db}

upgrader := Upgrader{}
upgrader.Register([]{upgrade})
upgrader.Run()
```
