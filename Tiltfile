docker_build('maxhorstmann/todo', '.', dockerfile = 'Dockerfile')

k8s_yaml('services.yaml')
k8s_resource('sql', port_forwards="1433")
#watch_file('pr-environment.yaml')

