#ifndef _SHELLCOMM_PARSE_
#define _SHELLCOMM_PARSE_

#include "b64/decode.h"
#include "nlohmann/json.hpp"

using json = nlohmann::json;

namespace ShellComm {

enum Mode {
  MODE_UNKNOWN = 0,
  MODE_WEBVIEW = 1
};

struct ParseResult {
  Mode mode;
  json json;
};

bool ParseArgv(int argc, const char **argv, ParseResult *result);

} // namespace ShellComm

#endif /* _SHELLCOMM_PARSE_ */
