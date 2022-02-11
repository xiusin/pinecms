export default {
  /**
   * 设置所选磁盘
   * @param state
   * @param disk
   */
  setDisk(state, disk) {
    state.selectedDisk = disk;
  },

  /**
   * 设置所选目录 
   * @param state
   * @param directory
   */
  setSelectedDirectory(state, directory) {
    state.selectedDirectory = directory;
  },
  /**
   * 所选目录的所有者
   * @param state 
   * @param  author 
   */
  setSelectedDirectoryOwner(state, author) {
    state.selectedDirectoryOwner = author;
  },
  /**
   * 添加历史记录 
   * @param state
   * @param path
   */
  addToHistory(state, path) {
    if (state.historyPointer < state.history.length - 1) {
      console.log("删除历史记录[", "historyPointer:", state.historyPointer, "history:", state.history);
      // 删除历史中的下一个元素
      state.history.splice(state.historyPointer + 1, Number.MAX_VALUE);
    }
    console.log("添加历史记录：", path);
    // 添加新路径
    state.history.push(path);
    // 改变 历史指针
    state.historyPointer += 1;
  },

  /**
   * 设置请求获取到的所选路径下的目录和文件 
   * @param state
   * @param data
   */
  setDirectoryContent(state, data, imageUrl = []) {
    state.directories = data.directories;
    state.files = data.files;
    state.imageUrl = imageUrl; // 原有的图片链接清空
    state.selectedDirectoryOwner = data.author;
  },
  /**
   * 设置图片链接数组
   * @param {*} state 
   * @param {Array} urlArr 
   */
  setImageUrl(state, urlArr) {
    state.imageUrl = urlArr;
  },

  /**
   * 清空之前所选的文件和文件夹
   * @param state
   */
  resetSelected(state) {
    state.selected.directories = [];
    state.selected.files = [];
  },

  /**
   * 重置文件排序配置 
   * @param state
   */
  resetSortSettings(state) {
    state.sort.field = 'name';
    state.sort.direction = 'up';
  },

  /**
   * 更改所选项目
   * @param state
   * @param type
   * @param path
   */
  changeSelected(state, {
    type,
    path
  }) {
    state.selected.directories = [];
    state.selected.files = [];
    state.selected[type].push(path);
  },

  /**
   * 文件是否全选
   * @param {*} state 
   * @param {Boolean} isCheckedAll 
   */
  setIsCheckedAll(state, isCheckedAll) {
    state.isCheckedAll = isCheckedAll;
  },
  /**
   * 设置全选效果属性
   * @param {*} state 
   * @param {Boolean} isIndeterminate 
   */
  setIsIndeterminate(state, isIndeterminate) {
    state.isIndeterminate = isIndeterminate;
  },
  /**
   * 设置上面两个属性
   * @param {*} state 
   * @param {Boolean} isIndeterminate 
   * @param {Boolean} isCheckedAll 
   */
  setChAndIn(state, {
    isIndeterminate,
    isCheckedAll
  }) {
    state.isIndeterminate = isIndeterminate;
    state.isCheckedAll = isCheckedAll;
    console.log("isIndeterminate:", isIndeterminate, "isCheckedAll:", isCheckedAll);
  },
  /**
   * 给复选框文件列表添加项
   * @param {*} state 
   * @param {Array|Object} checkedFile 
   */
  // setCheckedFiles(state, checkedFile, len = 0) {
  // if (Array.isArray(checkedFile)) {
  //   state.checkedFiles = checkedFile;
  // } else if (typeof checkedFile == "string") {
  //   state.checkedFiles[checkedFile] = len;
  // } else {
  //   state.checkedFiles.push(checkedFile)
  // }
  // console.log("设置选择的列表:", checkedFile, "选择的文件：", state.selected.directories, state.selected.files);
  // },
  /**
   * 删除复选框列表某一位置的元素
   * @param {*} state 
   * @param {Number} index 
   */
  spliceCheckedFiles(state, index) {
    state.checkedFiles.splice(index, 1);
  },
  /**
   * 设置所选项目 
   * @param state
   * @param type (directories, files)
   * @param path
   */
  setSelected(state, {
    type,
    path
  }) {
    state.selected[type].push(path);
  },
  /**
   * 全选
   * @param {*} state 
   * @param {Array} directives
   * @param {Array} files 
   */
  setAllSelected(state, {
    dir,
    file
  }) {
    if (dir) {
      dir = dir.map((item) => {
        return item.path
      })
      state.selected.directories.length = 0;
      state.selected.directories.splice(0, 0, ...dir);
    }
    if (file) {
      file = file.map((item) => {
        return item.path
      })
      state.selected.files.length = 0;
      state.selected.files.splice(0, 0, ...file);
    }
  },

  /**
   * 从数组中删除项 
   * @param state
   * @param arrayIndex
   */
  removeSelected(state, {
    type,
    path
  }) {
    const itemIndex = state.selected[type].indexOf(path);
    if (itemIndex !== -1) state.selected[type].splice(itemIndex, 1);
  },
  /**
   * 删除所有的选择项
   * @param {*} state 
   */
  removeAllSelected(state) {
    if (state.selected.directories.length) {
      state.selected.directories.splice(0, state.selected.directories.length);
    }
    if (state.selected.files.length) {
      state.selected.files.splice(0, state.selected.files.length);
    }
  },



  /**
   * 添加新文件
   * @param state
   * @param newFile
   */
  addNewFile(state, newFile) {
    state.files.push(newFile);
  },
  /**
   * 设置请求所有图片时当时排序的字段
   */
  setImagesSortField(state, {
    field,
    direction
  }) {
    state.imageUrl.sort = {
      field,
      direction
    };
  },
  /**
   * 更新文件
   * @param state
   * @param file
   */
  updateFile(state, file) {
    const itemIndex = state.files.findIndex((el) => el.basename === file.basename);
    if (itemIndex !== -1) state.files[itemIndex] = file;
  },

  /**
   * 添加新目录
   * @param state
   * @param newDirectory
   */
  addNewDirectory(state, newDirectory) {
    state.directories.push(newDirectory);
  },

  /**
   * 更改历史指针(back)
   * @param state
   */
  pointerBack(state) {
    state.historyPointer -= 1;
  },

  /**
   * 更改历史指针 (forward)
   * @param state
   */
  pointerForward(state) {
    state.historyPointer += 1;
  },



  /**
   * 重置历史记录
   * @param state
   */
  resetHistory(state) {
    state.history = [null];
    state.historyPointer = 0;
  },

  /**
   * 设置grid或table布局
   * Grid or Table
   * @param state
   * @param type
   */
  setView(state, type) {
    state.viewType = type;
  },

  /**
   * 设置排序配置-字段名称
   * @param state
   * @param field
   */
  setSortField(state, field) {
    state.sort.field = field;
  },

  /**
   * 设置排序配置-方向
   * @param state
   * @param direction
   */
  setSortDirection(state, direction) {
    state.sort.direction = direction;
  },



  /**
   * 按名称字段对表排序
   * @param state
   */
  sortByName(state) {
    if (state.sort.direction === 'up') {
      state.directories.sort((a, b) => a.basename.localeCompare(b.basename));
      state.files.sort((a, b) => a.basename.localeCompare(b.basename));
    } else {
      state.directories.sort((a, b) => b.basename.localeCompare(a.basename));
      state.files.sort((a, b) => b.basename.localeCompare(a.basename));
    }
  },

  /**
   * 按文件大小排序
   * @param state
   */
  sortBySize(state) {
    state.directories.sort((a, b) => a.basename.localeCompare(b.basename));

    if (state.sort.direction === 'up') {
      state.files.sort((a, b) => a.size - b.size);
    } else {
      state.files.sort((a, b) => b.size - a.size);
    }
  },

  /**
   * 按文件扩展名排序
   * @param state
   */
  sortByType(state) {
    state.directories.sort((a, b) => a.basename.localeCompare(b.basename));

    if (state.sort.direction === 'up') {
      state.files.sort((a, b) => a.extension.localeCompare(b.extension));
    } else {
      state.files.sort((a, b) => b.extension.localeCompare(a.extension));
    }
  },

  /**
   * 按日期排序
   * @param state
   */
  sortByDate(state) {
    if (state.sort.direction === 'up') {
      state.directories.sort((a, b) => a.timestamp - b.timestamp);
      state.files.sort((a, b) => a.timestamp - b.timestamp);
    } else {
      state.directories.sort((a, b) => b.timestamp - a.timestamp);
      state.files.sort((a, b) => b.timestamp - a.timestamp);
    }
  },
};