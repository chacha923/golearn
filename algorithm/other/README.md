[TOC]

## 贪心

> 直观地想一想，似乎不需要递归，只需要判断哪一个选择最具有「潜力」即可

不需要穷举的问题，可以考虑贪心算法


## 滑动窗口

### 模版
```
int left = 0, right = 0;

while (left < right && right < s.size()) {
    // 增大窗口
    window.add(s[right]);
    right++;
    
    while (window needs shrink) {
        // 缩小窗口
        window.remove(s[left]);
        left++;
    }
}
```
这个算法技巧的时间复杂度是 O(N)，比字符串暴力算法要高效得多。

简单说，指针 left, right 不会回退（它们的值只增不减），所以字符串/数组中的每个元素都只会进入窗口一次，然后被移出窗口一次，不会说有某些元素多次进入和离开窗口，所以算法的时间复杂度就和字符串/数组的长度成正比。
