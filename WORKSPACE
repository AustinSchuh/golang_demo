local_repository(
    name="io_bazel_rules_go",
    path="third_party/bazel-rules_go",
)

load("@io_bazel_rules_go//go:def.bzl", "go_download_sdk",
     "go_rules_dependencies", "go_register_toolchains")

go_download_sdk(
    name="go_sdk",
    sdks={
        "linux_amd64": (
            "go1.9.2.linux-amd64.tar.gz",
            "de874549d9a8d8d8062be05808509c09a88a248e77ec14eb77453530829ac02b",
        ),
        "darwin_amd64": (
            "go1.9.2.darwin-amd64.tar.gz",
            "73fd5840d55f5566d8db6c0ffdd187577e8ebe650c783f68bd27cbf95bde6743",
        ),
    },
    urls=["http://gerrit.peloton-tech.net/Build-Dependency-Files/{}", ], )
go_rules_dependencies()

go_register_toolchains()

load("@io_bazel_rules_go//go/private:repository_tools.bzl", "go_repository_tools")

go_repository_tools(
    name = "io_bazel_rules_go_repository_tools",
)

# Needed for io_bazel_rules_go tests
load("@io_bazel_rules_go//tests:bazel_tests.bzl", "test_environment")

test_environment()

new_http_archive(
  name = 'clang_3p6_repo',
  build_file = 'tools/cpp/clang_3p6/clang_3p6.BUILD',
  sha256 = '5ee9e04c55c2c99d0c0f83722102a49e98f485fc274f73111b33a7ac4e34e03e',
  url = 'http://frc971.org/Build-Dependencies/clang_3p6.tar.gz',
)
