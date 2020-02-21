package main

import (
	"errors"
	"fmt"
)

/*
1 数组的插入，删除，根据index查找，遍历
2 注意:index是下标，范围是[0,cap(array)-1]
3 数组中的数据为int类型
*/
type Array struct {
	data   []int //存储数组的数据
	length uint  //记录数组已经保存的元素个数
}

//根据指定容量，初始化数组
func NewArray(cap uint) *Array {
	if cap == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, cap, cap),
		length: 0,
	}
}

//返回数组已经保存的元素个数
func (this *Array) Len() uint {
	return this.length
}

//返回数组的长度（容量）
func (this *Array) Cap() uint {
	return uint(cap(this.data))
}

//判断index是否越界
func (this *Array) IsIndexOutOfRange(index uint) bool {
	if index < this.Cap()-1 {
		return false
	}
	return true
}

//根据索引查找数组的数据
func (this *Array) FindByIndex(index uint) (int, error) {
	if this.IsIndexOutOfRange(index) {
		return 0, errors.New("out of array range")
	}
	return this.data[index], nil
}

//根据index插入数据
//头部，中间，尾部插入
func (this *Array) InsertByIndex(index uint, data int) error {
	//判断数组是否已经满载
	if this.Len() == this.Cap() {
		return errors.New("full array")
	}
	//判断index是否越界
	if this.IsIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	var i uint
	for i = this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = data
	this.length++
	return nil
}

//根据index删除数据
func (this *Array) DeleteByIndex(index uint) (bool, error) {
	//判断index是否越界
	if this.IsIndexOutOfRange(index) {
		return false, errors.New("out of index range")
	}
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return true, nil
}

//遍历数组
func (this *Array) List() {
	fmt.Println(this.data)
}
