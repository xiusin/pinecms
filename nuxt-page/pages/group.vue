<template>
  <main>
    <div class="places">
      <p class="top">组团旅行</p>
      <h1>4人到夏威夷旅行</h1>
      <hr />
      <div class="wrapper">
        <div class="box item1"></div>
        <div class="box item2"></div>
        <div class="box item3"></div>
        <div class="box item4"></div>
        <div class="box item5"></div>
      </div>
    </div>

    <aside class="sidebar">
      <h3><icon-base icon-name="calendar"><icon-calendar /></icon-base> 日程</h3>
      <p class="top">星期天</p>
      <p>抵达：我们安顿下来并决定去浮潜。我们看到斑点的海豚，大量的鲸鱼 - 一条鲸鱼就在船底下，当然还有很多海龟。浮潜区域有点小，可能会变得拥挤，但我们仍然有很好的机会观看海龟。工作人员非常棒 - 杰森，杰克逊和谢恩。我们肯定会推荐这项活动，如果我们回到夏威夷，我们会再次这样做！</p>
      <p class="top">星期一</p>
      <p>经过如此激动人心的第一天，我们决定让它变得更轻松，并在我们的日常活动中享用美味的午餐。有一个美好的时光，有机会走几个街区的阿拉莫阿纳，并采取一些当地的食物，你不能真正回家。每一站都有一些新的和有趣的尝试，并没有一件事我尝试过，我真的不喜欢。</p>
    </aside>
  </main>
</template>

<style lang="scss" scoped>
.wrapper {
  margin: 10px 0 20px 0;
  width: 100%;
  display: grid;
  grid-gap: 10px;
  grid-template-columns: repeat(4, 24%);
  grid-template-rows: repeat(3, 120px);
  justify-content: center;
  align-content: end;
}

.box {
  border-radius: 3px;
}

$img-slug: hawaii !default;
$class-slug: item !default;

@for $i from 1 through 6 {
  .#{$class-slug}#{$i} {
    background: url('/hawaii#{$i}.jpg') center center no-repeat;
    background-size: cover;
  }
}

.item1 {
  grid-column: 1 / 5;
}

.item2 {
  grid-column: 1 / 3;
  grid-row: 2 / 4;
}

.item3 {
  grid-column: 3 / 5;
}

p {
  padding-bottom: 10px;
}
</style>
