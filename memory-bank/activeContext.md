# Active Context: aiyou-cli

## Current Work Focus

The current focus is on enhancing the usability and flexibility of the CLI by implementing two key features:

1. **Multi-line System Prompts in YAML Configuration**
   - Allow users to define complex system prompts with proper formatting
   - Support YAML's literal style syntax (using `|`) for multi-line text
   - Document the syntax in README and example configuration

2. **Standard Input Support for Messages**
   - Enable reading user messages from standard input (stdin)
   - Allow integration with pipes and other command-line tools
   - Implement with a new `-i/--stdin` flag

## Recent Changes

### Feature: Multi-line System Prompts
- Updated configuration handling to properly support YAML multi-line strings
- Added documentation and examples in README.md
- Updated example configuration file with multi-line prompt examples

### Feature: Standard Input Support
- Added new `-i/--stdin` flag to enable reading from stdin
- Implemented logic to read from stdin when flag is active
- Updated validation to accept either `-m/--message` or `-i/--stdin`
- Updated documentation in README.md

### Version Management
- Updated version from v0.1.0 to v1.1.0 to align with GitHub releases
- Created versioning.md to document the versioning and release process
- Implemented proper semantic versioning guidelines

### Documentation
- Updated README.md with new features
- Updated TODO.md to mark completed items
- Created comprehensive memory bank documentation

## Next Steps

### Short-term Tasks
- Test the new features with various input types and edge cases
- Consider adding unit tests for the new functionality
- Update the release version to reflect the new features

### Medium-term Improvements
- Consider adding support for reading system prompts from files
- Implement better error messages for common issues
- Add examples of integration with other command-line tools

### Long-term Vision
- Explore adding more interactive features
- Consider support for conversation history
- Investigate batch processing capabilities

## Active Decisions and Considerations

### Implementation Approach
- Decided to implement stdin support with a dedicated flag rather than auto-detection
- Chose to use YAML's built-in multi-line string support rather than custom parsing
- Prioritized backward compatibility to avoid breaking existing scripts

### User Experience
- Focused on making the new features intuitive and well-documented
- Ensured error messages are clear when using the new features incorrectly
- Maintained consistent command-line interface patterns

### Technical Considerations
- Ensured proper handling of newlines in multi-line prompts
- Addressed potential issues with stdin reading (EOF handling, etc.)
- Maintained clean separation of concerns in the codebase
