package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseModel(t *testing.T) {
	model := DedicatedDatabase{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", Name: "My Production Database", Version: "16", Specification: "s-2vcpu-2gb", Backend: "edge", Hostname: "db-myproject-mydb.fra.appwrite.center", ConnectionUser: "appwrite_user", ConnectionPassword: "••••••••", ConnectionString: "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require", Ssl: true, Status: "ready", ContainerStatus: "active", LifecycleState: "active", IdleTimeoutMinutes: 15, Cpu: 2000, Memory: 4096, Storage: 100, StorageStatus: "resizing", StorageTargetGb: 120, StorageClass: "ssd", StorageMaxGb: 100, NodePool: "db-pool-4vcpu-8gb", Replicas: 2, SyncMode: "async", NetworkMaxConnections: 500, NetworkIdleTimeoutSeconds: 900, NetworkIPAllowlist: []string{"test"}, BackupEnabled: true, Pitr: true, PitrRetentionDays: 14, StorageAutoscaling: true, StorageAutoscalingThresholdPercent: 85, StorageAutoscalingMaxGb: 30, MaintenanceWindowDay: "sun", MaintenanceWindowHourUtc: 3, MetricsEnabled: true, SqlApiEnabled: true, SqlApiAllowedStatements: []string{"test"}, SqlApiMaxRows: 10000, SqlApiMaxBytes: 10485760, SqlApiTimeoutSeconds: 30, Error: "string"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabase
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.ProjectId != model.ProjectId {
		t.Errorf("Expected ProjectId %v, got %v", model.ProjectId, result.ProjectId)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Version != model.Version {
		t.Errorf("Expected Version %v, got %v", model.Version, result.Version)
	}
	if result.Specification != model.Specification {
		t.Errorf("Expected Specification %v, got %v", model.Specification, result.Specification)
	}
	if result.Backend != model.Backend {
		t.Errorf("Expected Backend %v, got %v", model.Backend, result.Backend)
	}
	if result.Hostname != model.Hostname {
		t.Errorf("Expected Hostname %v, got %v", model.Hostname, result.Hostname)
	}
	if result.ConnectionUser != model.ConnectionUser {
		t.Errorf("Expected ConnectionUser %v, got %v", model.ConnectionUser, result.ConnectionUser)
	}
	if result.ConnectionPassword != model.ConnectionPassword {
		t.Errorf("Expected ConnectionPassword %v, got %v", model.ConnectionPassword, result.ConnectionPassword)
	}
	if result.ConnectionString != model.ConnectionString {
		t.Errorf("Expected ConnectionString %v, got %v", model.ConnectionString, result.ConnectionString)
	}
	if result.Ssl != model.Ssl {
		t.Errorf("Expected Ssl %v, got %v", model.Ssl, result.Ssl)
	}
	if result.Status != model.Status {
		t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
	}
	if result.ContainerStatus != model.ContainerStatus {
		t.Errorf("Expected ContainerStatus %v, got %v", model.ContainerStatus, result.ContainerStatus)
	}
	if result.LifecycleState != model.LifecycleState {
		t.Errorf("Expected LifecycleState %v, got %v", model.LifecycleState, result.LifecycleState)
	}
	if result.IdleTimeoutMinutes != model.IdleTimeoutMinutes {
		t.Errorf("Expected IdleTimeoutMinutes %v, got %v", model.IdleTimeoutMinutes, result.IdleTimeoutMinutes)
	}
	if result.Cpu != model.Cpu {
		t.Errorf("Expected Cpu %v, got %v", model.Cpu, result.Cpu)
	}
	if result.Memory != model.Memory {
		t.Errorf("Expected Memory %v, got %v", model.Memory, result.Memory)
	}
	if result.Storage != model.Storage {
		t.Errorf("Expected Storage %v, got %v", model.Storage, result.Storage)
	}
	if result.StorageStatus != model.StorageStatus {
		t.Errorf("Expected StorageStatus %v, got %v", model.StorageStatus, result.StorageStatus)
	}
	if result.StorageTargetGb != model.StorageTargetGb {
		t.Errorf("Expected StorageTargetGb %v, got %v", model.StorageTargetGb, result.StorageTargetGb)
	}
	if result.StorageClass != model.StorageClass {
		t.Errorf("Expected StorageClass %v, got %v", model.StorageClass, result.StorageClass)
	}
	if result.StorageMaxGb != model.StorageMaxGb {
		t.Errorf("Expected StorageMaxGb %v, got %v", model.StorageMaxGb, result.StorageMaxGb)
	}
	if result.NodePool != model.NodePool {
		t.Errorf("Expected NodePool %v, got %v", model.NodePool, result.NodePool)
	}
	if result.Replicas != model.Replicas {
		t.Errorf("Expected Replicas %v, got %v", model.Replicas, result.Replicas)
	}
	if result.SyncMode != model.SyncMode {
		t.Errorf("Expected SyncMode %v, got %v", model.SyncMode, result.SyncMode)
	}
	if result.NetworkMaxConnections != model.NetworkMaxConnections {
		t.Errorf("Expected NetworkMaxConnections %v, got %v", model.NetworkMaxConnections, result.NetworkMaxConnections)
	}
	if result.NetworkIdleTimeoutSeconds != model.NetworkIdleTimeoutSeconds {
		t.Errorf("Expected NetworkIdleTimeoutSeconds %v, got %v", model.NetworkIdleTimeoutSeconds, result.NetworkIdleTimeoutSeconds)
	}
	if result.BackupEnabled != model.BackupEnabled {
		t.Errorf("Expected BackupEnabled %v, got %v", model.BackupEnabled, result.BackupEnabled)
	}
	if result.Pitr != model.Pitr {
		t.Errorf("Expected Pitr %v, got %v", model.Pitr, result.Pitr)
	}
	if result.PitrRetentionDays != model.PitrRetentionDays {
		t.Errorf("Expected PitrRetentionDays %v, got %v", model.PitrRetentionDays, result.PitrRetentionDays)
	}
	if result.StorageAutoscaling != model.StorageAutoscaling {
		t.Errorf("Expected StorageAutoscaling %v, got %v", model.StorageAutoscaling, result.StorageAutoscaling)
	}
	if result.StorageAutoscalingThresholdPercent != model.StorageAutoscalingThresholdPercent {
		t.Errorf("Expected StorageAutoscalingThresholdPercent %v, got %v", model.StorageAutoscalingThresholdPercent, result.StorageAutoscalingThresholdPercent)
	}
	if result.StorageAutoscalingMaxGb != model.StorageAutoscalingMaxGb {
		t.Errorf("Expected StorageAutoscalingMaxGb %v, got %v", model.StorageAutoscalingMaxGb, result.StorageAutoscalingMaxGb)
	}
	if result.MaintenanceWindowDay != model.MaintenanceWindowDay {
		t.Errorf("Expected MaintenanceWindowDay %v, got %v", model.MaintenanceWindowDay, result.MaintenanceWindowDay)
	}
	if result.MaintenanceWindowHourUtc != model.MaintenanceWindowHourUtc {
		t.Errorf("Expected MaintenanceWindowHourUtc %v, got %v", model.MaintenanceWindowHourUtc, result.MaintenanceWindowHourUtc)
	}
	if result.MetricsEnabled != model.MetricsEnabled {
		t.Errorf("Expected MetricsEnabled %v, got %v", model.MetricsEnabled, result.MetricsEnabled)
	}
	if result.SqlApiEnabled != model.SqlApiEnabled {
		t.Errorf("Expected SqlApiEnabled %v, got %v", model.SqlApiEnabled, result.SqlApiEnabled)
	}
	if result.SqlApiMaxRows != model.SqlApiMaxRows {
		t.Errorf("Expected SqlApiMaxRows %v, got %v", model.SqlApiMaxRows, result.SqlApiMaxRows)
	}
	if result.SqlApiMaxBytes != model.SqlApiMaxBytes {
		t.Errorf("Expected SqlApiMaxBytes %v, got %v", model.SqlApiMaxBytes, result.SqlApiMaxBytes)
	}
	if result.SqlApiTimeoutSeconds != model.SqlApiTimeoutSeconds {
		t.Errorf("Expected SqlApiTimeoutSeconds %v, got %v", model.SqlApiTimeoutSeconds, result.SqlApiTimeoutSeconds)
	}
	if result.Error != model.Error {
		t.Errorf("Expected Error %v, got %v", model.Error, result.Error)
	}
}

func TestDedicatedDatabaseModelDefaults(t *testing.T) {
	var fresh DedicatedDatabase
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Id)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $id to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CreatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $createdAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UpdatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $updatedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProjectId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing projectId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing name to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Version)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing version to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Specification)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing specification to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Backend)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing backend to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Hostname)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing hostname to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ConnectionUser)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing connectionUser to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ConnectionPassword)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing connectionPassword to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ConnectionString)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing connectionString to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Ssl)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ssl to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"provisioning\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Status)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing status to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ContainerStatus)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing containerStatus to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LifecycleState)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing lifecycleState to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("15"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.IdleTimeoutMinutes)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing idleTimeoutMinutes to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("500"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Cpu)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing cpu to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("512"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Memory)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing memory to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("10"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Storage)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storage to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"idle\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageStatus)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageStatus to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageTargetGb)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageTargetGb to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"ssd\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageClass)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageClass to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageMaxGb)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageMaxGb to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.NodePool)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing nodePool to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Replicas)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing replicas to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SyncMode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing syncMode to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("100"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.NetworkMaxConnections)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing networkMaxConnections to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("900"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.NetworkIdleTimeoutSeconds)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing networkIdleTimeoutSeconds to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.NetworkIPAllowlist)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing networkIPAllowlist to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BackupEnabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing backupEnabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Pitr)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing pitr to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("7"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PitrRetentionDays)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing pitrRetentionDays to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageAutoscaling)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageAutoscaling to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("85"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageAutoscalingThresholdPercent)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageAutoscalingThresholdPercent to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageAutoscalingMaxGb)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageAutoscalingMaxGb to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"sun\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.MaintenanceWindowDay)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing maintenanceWindowDay to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("3"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.MaintenanceWindowHourUtc)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing maintenanceWindowHourUtc to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.MetricsEnabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing metricsEnabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SqlApiEnabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing sqlApiEnabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[\"SELECT\",\"INSERT\",\"UPDATE\",\"DELETE\"]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SqlApiAllowedStatements)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing sqlApiAllowedStatements to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("10000"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SqlApiMaxRows)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing sqlApiMaxRows to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("10485760"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SqlApiMaxBytes)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing sqlApiMaxBytes to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("30"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SqlApiTimeoutSeconds)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing sqlApiTimeoutSeconds to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Error)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing error to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabase{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", Name: "My Production Database", Version: "16", Specification: "s-2vcpu-2gb", Backend: "edge", Hostname: "db-myproject-mydb.fra.appwrite.center", ConnectionUser: "appwrite_user", ConnectionPassword: "••••••••", ConnectionString: "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require", Ssl: true, Status: "ready", ContainerStatus: "active", LifecycleState: "active", IdleTimeoutMinutes: 15, Cpu: 2000, Memory: 4096, Storage: 100, StorageStatus: "resizing", StorageTargetGb: 120, StorageClass: "ssd", StorageMaxGb: 100, NodePool: "db-pool-4vcpu-8gb", Replicas: 2, SyncMode: "async", NetworkMaxConnections: 500, NetworkIdleTimeoutSeconds: 900, NetworkIPAllowlist: []string{"test"}, BackupEnabled: true, Pitr: true, PitrRetentionDays: 14, StorageAutoscaling: true, StorageAutoscalingThresholdPercent: 85, StorageAutoscalingMaxGb: 30, MaintenanceWindowDay: "sun", MaintenanceWindowHourUtc: 3, MetricsEnabled: true, SqlApiEnabled: true, SqlApiAllowedStatements: []string{"test"}, SqlApiMaxRows: 10000, SqlApiMaxBytes: 10485760, SqlApiTimeoutSeconds: 30, Error: "string"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
