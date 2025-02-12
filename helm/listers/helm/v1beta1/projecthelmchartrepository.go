// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	helmv1beta1 "github.com/openshift/api/helm/v1beta1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ProjectHelmChartRepositoryLister helps list ProjectHelmChartRepositories.
// All objects returned here must be treated as read-only.
type ProjectHelmChartRepositoryLister interface {
	// List lists all ProjectHelmChartRepositories in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*helmv1beta1.ProjectHelmChartRepository, err error)
	// ProjectHelmChartRepositories returns an object that can list and get ProjectHelmChartRepositories.
	ProjectHelmChartRepositories(namespace string) ProjectHelmChartRepositoryNamespaceLister
	ProjectHelmChartRepositoryListerExpansion
}

// projectHelmChartRepositoryLister implements the ProjectHelmChartRepositoryLister interface.
type projectHelmChartRepositoryLister struct {
	listers.ResourceIndexer[*helmv1beta1.ProjectHelmChartRepository]
}

// NewProjectHelmChartRepositoryLister returns a new ProjectHelmChartRepositoryLister.
func NewProjectHelmChartRepositoryLister(indexer cache.Indexer) ProjectHelmChartRepositoryLister {
	return &projectHelmChartRepositoryLister{listers.New[*helmv1beta1.ProjectHelmChartRepository](indexer, helmv1beta1.Resource("projecthelmchartrepository"))}
}

// ProjectHelmChartRepositories returns an object that can list and get ProjectHelmChartRepositories.
func (s *projectHelmChartRepositoryLister) ProjectHelmChartRepositories(namespace string) ProjectHelmChartRepositoryNamespaceLister {
	return projectHelmChartRepositoryNamespaceLister{listers.NewNamespaced[*helmv1beta1.ProjectHelmChartRepository](s.ResourceIndexer, namespace)}
}

// ProjectHelmChartRepositoryNamespaceLister helps list and get ProjectHelmChartRepositories.
// All objects returned here must be treated as read-only.
type ProjectHelmChartRepositoryNamespaceLister interface {
	// List lists all ProjectHelmChartRepositories in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*helmv1beta1.ProjectHelmChartRepository, err error)
	// Get retrieves the ProjectHelmChartRepository from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*helmv1beta1.ProjectHelmChartRepository, error)
	ProjectHelmChartRepositoryNamespaceListerExpansion
}

// projectHelmChartRepositoryNamespaceLister implements the ProjectHelmChartRepositoryNamespaceLister
// interface.
type projectHelmChartRepositoryNamespaceLister struct {
	listers.ResourceIndexer[*helmv1beta1.ProjectHelmChartRepository]
}
