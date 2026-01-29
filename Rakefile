# frozen_string_literal: true

# Copyright (c) 2025 kk
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

require 'bundler/setup'
require 'kk/git/rake_tasks'

task default: %w[push]

task :push do
  Rake::Task['git:auto_commit_push'].invoke
  Rake::Task['tag'].invoke
end

desc 'Create and push a random version tag (v0.0.x)'
task :tag do
  version = "v0.0.#{rand(1000..9999)}"
  puts "Creating tag: #{version}"
  sh "git tag #{version}"
  sh "git push origin #{version}"
end