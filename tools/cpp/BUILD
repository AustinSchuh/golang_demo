package(default_visibility = ["//visibility:public"])

cc_library(
    name = "empty_main",
    srcs = ["empty_main.c"],
)

cc_library(
    name = "malloc",
    deps = [],
)

cc_library(
    name = "stl",
)

filegroup(
    name = "empty",
    srcs = [],
)

# This is the entry point for --crosstool_top.
cc_toolchain_suite(
    name = "toolchain",
    toolchains = {
        "k8|clang": ":cc-compiler-k8",
    },
)

# Compiler inputs given by --copt etc in //tools:bazel.rc.
filegroup(
    name = "flags_compiler_inputs",
    srcs = [],
)

filegroup(
    name = "clang_3p6_all_files",
    srcs = [
        ":flags_compiler_inputs",
        "//tools/cpp/clang_3p6:as",
        "//tools/cpp/clang_3p6:tool-wrappers",
        "@clang_3p6_repo//:compiler_pieces",
    ],
)

filegroup(
    name = "clang_3p6_linker_files",
    srcs = [
        "//tools/cpp/clang_3p6:clang-symlinks",
        "//tools/cpp/clang_3p6:ld",
        "//tools/cpp/clang_3p6:ar",
        "//tools/cpp/clang_3p6:clang",
        #'//tools/cpp/clang_3p6:libs',
        "@clang_3p6_repo//:compiler_pieces",
    ],
)

filegroup(
    name = "clang_3p6_compiler_files",
    srcs = [
        "//tools/cpp/clang_3p6:clang",
        "//tools/cpp/clang_3p6:ld",
        "@clang_3p6_repo//:compiler_components",
        "@clang_3p6_repo//:compiler_pieces",
    ],
)

cc_toolchain(
    name = "cc-compiler-k8",
    all_files = ":clang_3p6_all_files",
    compiler_files = ":clang_3p6_compiler_files",
    cpu = "k8",
    dwp_files = ":empty",
    dynamic_runtime_libs = [":empty"],
    linker_files = ":clang_3p6_linker_files",
    objcopy_files = "//tools/cpp/clang_3p6:objcopy",
    static_runtime_libs = [":empty"],
    strip_files = "//tools/cpp/clang_3p6:strip",
    supports_param_files = 1,
)
