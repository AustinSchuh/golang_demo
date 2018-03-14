package(default_visibility = ['//tools/cpp:__pkg__'])

cc_library(
  name = "libpthread",
  srcs = [
    "clang_more_libs/libpthread.so",
  ],
  visibility = ["//visibility:public"],
)

filegroup(
  name = 'ar',
  srcs = [
    '@clang_3p6_repo//:ar',
    'x86_64-linux-gnu-ar',
  ],
)

filegroup(
  name = 'ld',
  srcs = [
    '@clang_3p6_repo//:ld',
    'x86_64-linux-gnu-ld',
  ],
)

filegroup(
  name = 'nm',
  srcs = [
    '@clang_3p6_repo//:nm',
    'x86_64-linux-gnu-nm',
  ],
)

filegroup(
  name = 'objcopy',
  srcs = [
    '@clang_3p6_repo//:objcopy',
    'x86_64-linux-gnu-objcopy',
  ],
)

filegroup(
  name = 'objdump',
  srcs = [
    '@clang_3p6_repo//:objdump',
    'x86_64-linux-gnu-objdump',
  ],
)

filegroup(
  name = 'strip',
  srcs = [
    '@clang_3p6_repo//:strip',
    'x86_64-linux-gnu-strip',
  ],
)

filegroup(
  name = 'as',
  srcs = [
    '@clang_3p6_repo//:as',
    'x86_64-linux-gnu-as',
  ],
)

filegroup(
  name = 'clang',
  srcs = [
    '@clang_3p6_repo//:clang',
    'x86_64-linux-gnu-clang-3.6',
  ],
)

filegroup(
  name = 'tool-wrappers',
  srcs = [
    ':clang',
    ':ar',
    ':ld',
    ':nm',
    ':objcopy',
    ':objdump',
    ':strip',
    ':as',
  ],
)

filegroup(
    name = "clang-symlinks",
    srcs = glob([
        "clang_more_libs/**",
    ]),
)
