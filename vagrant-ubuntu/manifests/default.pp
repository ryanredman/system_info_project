exec { 'apt-update':
    command => '/usr/bin/apt-get update'
}

package { 'git':
    require => Exec['apt-update'],
    ensure => installed,
}

package { 'golang':
    require => Exec['apt-update'],
    ensure => installed,
}

exec { 'fetch_repo':
    command => '/usr/bin/git clone https://github.com/ryanredman/system_info_project.git',
    creates => '/home/vagrant/system_info_project',
    require => Package['golang'],
}

exec { 'install_system_info':
    environment => [ 'GOPATH=/home/vagrant/system_info_project/system_info', 
                'GOBIN=/home/vagrant/system_info_project/system_info/bin' ],
    command => '/usr/bin/go install system_info',
    require => Exec['fetch_repo'],
}

exec { 'run_system_info':
    unless => '/usr/bin/pgrep system_info',
    command => '/home/vagrant/system_info_project/system_info/bin/system_info &',
    require => Exec['install_system_info'],
}
