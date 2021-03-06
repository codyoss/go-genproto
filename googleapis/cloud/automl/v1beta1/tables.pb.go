// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/tables.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Metadata for a dataset used for AutoML Tables.
type TablesDatasetMetadata struct {
	// Output only. The table_spec_id of the primary table of this dataset.
	PrimaryTableSpecId string `protobuf:"bytes,1,opt,name=primary_table_spec_id,json=primaryTableSpecId,proto3" json:"primary_table_spec_id,omitempty"`
	// column_spec_id of the primary table's column that should be used as the
	// training & prediction target.
	// This column must be non-nullable and have one of following data types
	// (otherwise model creation will error):
	//
	// * CATEGORY
	//
	// * FLOAT64
	//
	// If the type is CATEGORY , only up to
	// 100 unique values may exist in that column across all rows.
	//
	// NOTE: Updates of this field will instantly affect any other users
	// concurrently working with the dataset.
	TargetColumnSpecId string `protobuf:"bytes,2,opt,name=target_column_spec_id,json=targetColumnSpecId,proto3" json:"target_column_spec_id,omitempty"`
	// column_spec_id of the primary table's column that should be used as the
	// weight column, i.e. the higher the value the more important the row will be
	// during model training.
	// Required type: FLOAT64.
	// Allowed values: 0 to 10000, inclusive on both ends; 0 means the row is
	//                 ignored for training.
	// If not set all rows are assumed to have equal weight of 1.
	// NOTE: Updates of this field will instantly affect any other users
	// concurrently working with the dataset.
	WeightColumnSpecId string `protobuf:"bytes,3,opt,name=weight_column_spec_id,json=weightColumnSpecId,proto3" json:"weight_column_spec_id,omitempty"`
	// column_spec_id of the primary table column which specifies a possible ML
	// use of the row, i.e. the column will be used to split the rows into TRAIN,
	// VALIDATE and TEST sets.
	// Required type: STRING.
	// This column, if set, must either have all of `TRAIN`, `VALIDATE`, `TEST`
	// among its values, or only have `TEST`, `UNASSIGNED` values. In the latter
	// case the rows with `UNASSIGNED` value will be assigned by AutoML. Note
	// that if a given ml use distribution makes it impossible to create a "good"
	// model, that call will error describing the issue.
	// If both this column_spec_id and primary table's time_column_spec_id are not
	// set, then all rows are treated as `UNASSIGNED`.
	// NOTE: Updates of this field will instantly affect any other users
	// concurrently working with the dataset.
	MlUseColumnSpecId string `protobuf:"bytes,4,opt,name=ml_use_column_spec_id,json=mlUseColumnSpecId,proto3" json:"ml_use_column_spec_id,omitempty"`
	// Output only. Correlations between
	//
	// [TablesDatasetMetadata.target_column_spec_id][google.cloud.automl.v1beta1.TablesDatasetMetadata.target_column_spec_id],
	// and other columns of the
	//
	// [TablesDatasetMetadataprimary_table][google.cloud.automl.v1beta1.TablesDatasetMetadata.primary_table_spec_id].
	// Only set if the target column is set. Mapping from other column spec id to
	// its CorrelationStats with the target column.
	// This field may be stale, see the stats_update_time field for
	// for the timestamp at which these stats were last updated.
	TargetColumnCorrelations map[string]*CorrelationStats `protobuf:"bytes,6,rep,name=target_column_correlations,json=targetColumnCorrelations,proto3" json:"target_column_correlations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. The most recent timestamp when target_column_correlations
	// field and all descendant ColumnSpec.data_stats and
	// ColumnSpec.top_correlated_columns fields were last (re-)generated. Any
	// changes that happened to the dataset afterwards are not reflected in these
	// fields values. The regeneration happens in the background on a best effort
	// basis.
	StatsUpdateTime      *timestamp.Timestamp `protobuf:"bytes,7,opt,name=stats_update_time,json=statsUpdateTime,proto3" json:"stats_update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TablesDatasetMetadata) Reset()         { *m = TablesDatasetMetadata{} }
func (m *TablesDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TablesDatasetMetadata) ProtoMessage()    {}
func (*TablesDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4edb25480826735, []int{0}
}

func (m *TablesDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesDatasetMetadata.Unmarshal(m, b)
}
func (m *TablesDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TablesDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesDatasetMetadata.Merge(m, src)
}
func (m *TablesDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TablesDatasetMetadata.Size(m)
}
func (m *TablesDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TablesDatasetMetadata proto.InternalMessageInfo

func (m *TablesDatasetMetadata) GetPrimaryTableSpecId() string {
	if m != nil {
		return m.PrimaryTableSpecId
	}
	return ""
}

func (m *TablesDatasetMetadata) GetTargetColumnSpecId() string {
	if m != nil {
		return m.TargetColumnSpecId
	}
	return ""
}

func (m *TablesDatasetMetadata) GetWeightColumnSpecId() string {
	if m != nil {
		return m.WeightColumnSpecId
	}
	return ""
}

func (m *TablesDatasetMetadata) GetMlUseColumnSpecId() string {
	if m != nil {
		return m.MlUseColumnSpecId
	}
	return ""
}

func (m *TablesDatasetMetadata) GetTargetColumnCorrelations() map[string]*CorrelationStats {
	if m != nil {
		return m.TargetColumnCorrelations
	}
	return nil
}

func (m *TablesDatasetMetadata) GetStatsUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.StatsUpdateTime
	}
	return nil
}

// Model metadata specific to AutoML Tables.
type TablesModelMetadata struct {
	// Additional optimization objective configuration. Required for
	// `MAXIMIZE_PRECISION_AT_RECALL` and `MAXIMIZE_RECALL_AT_PRECISION`,
	// otherwise unused.
	//
	// Types that are valid to be assigned to AdditionalOptimizationObjectiveConfig:
	//	*TablesModelMetadata_OptimizationObjectiveRecallValue
	//	*TablesModelMetadata_OptimizationObjectivePrecisionValue
	AdditionalOptimizationObjectiveConfig isTablesModelMetadata_AdditionalOptimizationObjectiveConfig `protobuf_oneof:"additional_optimization_objective_config"`
	// Column spec of the dataset's primary table's column the model is
	// predicting. Snapshotted when model creation started.
	// Only 3 fields are used:
	// name - May be set on CreateModel, if it's not then the ColumnSpec
	//        corresponding to the current target_column_spec_id of the dataset
	//        the model is trained from is used.
	//        If neither is set, CreateModel will error.
	// display_name - Output only.
	// data_type - Output only.
	TargetColumnSpec *ColumnSpec `protobuf:"bytes,2,opt,name=target_column_spec,json=targetColumnSpec,proto3" json:"target_column_spec,omitempty"`
	// Column specs of the dataset's primary table's columns, on which
	// the model is trained and which are used as the input for predictions.
	// The
	//
	// [target_column][google.cloud.automl.v1beta1.TablesModelMetadata.target_column_spec]
	// as well as, according to dataset's state upon model creation,
	//
	// [weight_column][google.cloud.automl.v1beta1.TablesDatasetMetadata.weight_column_spec_id],
	// and
	//
	// [ml_use_column][google.cloud.automl.v1beta1.TablesDatasetMetadata.ml_use_column_spec_id]
	// must never be included here.
	//
	// Only 3 fields are used:
	//
	// * name - May be set on CreateModel, if set only the columns specified are
	//   used, otherwise all primary table's columns (except the ones listed
	//   above) are used for the training and prediction input.
	//
	// * display_name - Output only.
	//
	// * data_type - Output only.
	InputFeatureColumnSpecs []*ColumnSpec `protobuf:"bytes,3,rep,name=input_feature_column_specs,json=inputFeatureColumnSpecs,proto3" json:"input_feature_column_specs,omitempty"`
	// Objective function the model is optimizing towards. The training process
	// creates a model that maximizes/minimizes the value of the objective
	// function over the validation set.
	//
	// The supported optimization objectives depend on the prediction type.
	// If the field is not set, a default objective function is used.
	//
	// CLASSIFICATION_BINARY:
	//   "MAXIMIZE_AU_ROC" (default) - Maximize the area under the receiver
	//                                 operating characteristic (ROC) curve.
	//   "MINIMIZE_LOG_LOSS" - Minimize log loss.
	//   "MAXIMIZE_AU_PRC" - Maximize the area under the precision-recall curve.
	//   "MAXIMIZE_PRECISION_AT_RECALL" - Maximize precision for a specified
	//                                   recall value.
	//   "MAXIMIZE_RECALL_AT_PRECISION" - Maximize recall for a specified
	//                                    precision value.
	//
	// CLASSIFICATION_MULTI_CLASS :
	//   "MINIMIZE_LOG_LOSS" (default) - Minimize log loss.
	//
	//
	// REGRESSION:
	//   "MINIMIZE_RMSE" (default) - Minimize root-mean-squared error (RMSE).
	//   "MINIMIZE_MAE" - Minimize mean-absolute error (MAE).
	//   "MINIMIZE_RMSLE" - Minimize root-mean-squared log error (RMSLE).
	OptimizationObjective string `protobuf:"bytes,4,opt,name=optimization_objective,json=optimizationObjective,proto3" json:"optimization_objective,omitempty"`
	// Output only. Auxiliary information for each of the
	// input_feature_column_specs with respect to this particular model.
	TablesModelColumnInfo []*TablesModelColumnInfo `protobuf:"bytes,5,rep,name=tables_model_column_info,json=tablesModelColumnInfo,proto3" json:"tables_model_column_info,omitempty"`
	// Required. The train budget of creating this model, expressed in milli node
	// hours i.e. 1,000 value in this field means 1 node hour.
	//
	// The training cost of the model will not exceed this budget. The final cost
	// will be attempted to be close to the budget, though may end up being (even)
	// noticeably smaller - at the backend's discretion. This especially may
	// happen when further model training ceases to provide any improvements.
	//
	// If the budget is set to a value known to be insufficient to train a
	// model for the given dataset, the training won't be attempted and
	// will error.
	//
	// The train budget must be between 1,000 and 72,000 milli node hours,
	// inclusive.
	TrainBudgetMilliNodeHours int64 `protobuf:"varint,6,opt,name=train_budget_milli_node_hours,json=trainBudgetMilliNodeHours,proto3" json:"train_budget_milli_node_hours,omitempty"`
	// Output only. The actual training cost of the model, expressed in milli
	// node hours, i.e. 1,000 value in this field means 1 node hour. Guaranteed
	// to not exceed the train budget.
	TrainCostMilliNodeHours int64 `protobuf:"varint,7,opt,name=train_cost_milli_node_hours,json=trainCostMilliNodeHours,proto3" json:"train_cost_milli_node_hours,omitempty"`
	// Use the entire training budget. This disables the early stopping feature.
	// By default, the early stopping feature is enabled, which means that AutoML
	// Tables might stop training before the entire training budget has been used.
	DisableEarlyStopping bool     `protobuf:"varint,12,opt,name=disable_early_stopping,json=disableEarlyStopping,proto3" json:"disable_early_stopping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TablesModelMetadata) Reset()         { *m = TablesModelMetadata{} }
func (m *TablesModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TablesModelMetadata) ProtoMessage()    {}
func (*TablesModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4edb25480826735, []int{1}
}

func (m *TablesModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesModelMetadata.Unmarshal(m, b)
}
func (m *TablesModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesModelMetadata.Marshal(b, m, deterministic)
}
func (m *TablesModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesModelMetadata.Merge(m, src)
}
func (m *TablesModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TablesModelMetadata.Size(m)
}
func (m *TablesModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TablesModelMetadata proto.InternalMessageInfo

type isTablesModelMetadata_AdditionalOptimizationObjectiveConfig interface {
	isTablesModelMetadata_AdditionalOptimizationObjectiveConfig()
}

type TablesModelMetadata_OptimizationObjectiveRecallValue struct {
	OptimizationObjectiveRecallValue float32 `protobuf:"fixed32,17,opt,name=optimization_objective_recall_value,json=optimizationObjectiveRecallValue,proto3,oneof"`
}

type TablesModelMetadata_OptimizationObjectivePrecisionValue struct {
	OptimizationObjectivePrecisionValue float32 `protobuf:"fixed32,18,opt,name=optimization_objective_precision_value,json=optimizationObjectivePrecisionValue,proto3,oneof"`
}

func (*TablesModelMetadata_OptimizationObjectiveRecallValue) isTablesModelMetadata_AdditionalOptimizationObjectiveConfig() {
}

func (*TablesModelMetadata_OptimizationObjectivePrecisionValue) isTablesModelMetadata_AdditionalOptimizationObjectiveConfig() {
}

func (m *TablesModelMetadata) GetAdditionalOptimizationObjectiveConfig() isTablesModelMetadata_AdditionalOptimizationObjectiveConfig {
	if m != nil {
		return m.AdditionalOptimizationObjectiveConfig
	}
	return nil
}

func (m *TablesModelMetadata) GetOptimizationObjectiveRecallValue() float32 {
	if x, ok := m.GetAdditionalOptimizationObjectiveConfig().(*TablesModelMetadata_OptimizationObjectiveRecallValue); ok {
		return x.OptimizationObjectiveRecallValue
	}
	return 0
}

func (m *TablesModelMetadata) GetOptimizationObjectivePrecisionValue() float32 {
	if x, ok := m.GetAdditionalOptimizationObjectiveConfig().(*TablesModelMetadata_OptimizationObjectivePrecisionValue); ok {
		return x.OptimizationObjectivePrecisionValue
	}
	return 0
}

func (m *TablesModelMetadata) GetTargetColumnSpec() *ColumnSpec {
	if m != nil {
		return m.TargetColumnSpec
	}
	return nil
}

func (m *TablesModelMetadata) GetInputFeatureColumnSpecs() []*ColumnSpec {
	if m != nil {
		return m.InputFeatureColumnSpecs
	}
	return nil
}

func (m *TablesModelMetadata) GetOptimizationObjective() string {
	if m != nil {
		return m.OptimizationObjective
	}
	return ""
}

func (m *TablesModelMetadata) GetTablesModelColumnInfo() []*TablesModelColumnInfo {
	if m != nil {
		return m.TablesModelColumnInfo
	}
	return nil
}

func (m *TablesModelMetadata) GetTrainBudgetMilliNodeHours() int64 {
	if m != nil {
		return m.TrainBudgetMilliNodeHours
	}
	return 0
}

func (m *TablesModelMetadata) GetTrainCostMilliNodeHours() int64 {
	if m != nil {
		return m.TrainCostMilliNodeHours
	}
	return 0
}

func (m *TablesModelMetadata) GetDisableEarlyStopping() bool {
	if m != nil {
		return m.DisableEarlyStopping
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TablesModelMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TablesModelMetadata_OptimizationObjectiveRecallValue)(nil),
		(*TablesModelMetadata_OptimizationObjectivePrecisionValue)(nil),
	}
}

// Contains annotation details specific to Tables.
type TablesAnnotation struct {
	// Output only. A confidence estimate between 0.0 and 1.0, inclusive. A higher
	// value means greater confidence in the returned value.
	// For
	//
	// [target_column_spec][google.cloud.automl.v1beta1.TablesModelMetadata.target_column_spec]
	// of FLOAT64 data type the score is not populated.
	Score float32 `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
	// Output only. Only populated when
	//
	// [target_column_spec][google.cloud.automl.v1beta1.TablesModelMetadata.target_column_spec]
	// has FLOAT64 data type. An interval in which the exactly correct target
	// value has 95% chance to be in.
	PredictionInterval *DoubleRange `protobuf:"bytes,4,opt,name=prediction_interval,json=predictionInterval,proto3" json:"prediction_interval,omitempty"`
	// The predicted value of the row's
	//
	// [target_column][google.cloud.automl.v1beta1.TablesModelMetadata.target_column_spec].
	// The value depends on the column's DataType:
	//
	// * CATEGORY - the predicted (with the above confidence `score`) CATEGORY
	//   value.
	//
	// * FLOAT64 - the predicted (with above `prediction_interval`) FLOAT64 value.
	Value *_struct.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Output only. Auxiliary information for each of the model's
	//
	// [input_feature_column_specs][google.cloud.automl.v1beta1.TablesModelMetadata.input_feature_column_specs]
	// with respect to this particular prediction.
	// If no other fields than
	//
	// [column_spec_name][google.cloud.automl.v1beta1.TablesModelColumnInfo.column_spec_name]
	// and
	//
	// [column_display_name][google.cloud.automl.v1beta1.TablesModelColumnInfo.column_display_name]
	// would be populated, then this whole field is not.
	TablesModelColumnInfo []*TablesModelColumnInfo `protobuf:"bytes,3,rep,name=tables_model_column_info,json=tablesModelColumnInfo,proto3" json:"tables_model_column_info,omitempty"`
	// Output only. Stores the prediction score for the baseline example, which
	// is defined as the example with all values set to their baseline values.
	// This is used as part of the Sampled Shapley explanation of the model's
	// prediction. This field is populated only when feature importance is
	// requested. For regression models, this holds the baseline prediction for
	// the baseline example. For classification models, this holds the baseline
	// prediction for the baseline example for the argmax class.
	BaselineScore        float32  `protobuf:"fixed32,5,opt,name=baseline_score,json=baselineScore,proto3" json:"baseline_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TablesAnnotation) Reset()         { *m = TablesAnnotation{} }
func (m *TablesAnnotation) String() string { return proto.CompactTextString(m) }
func (*TablesAnnotation) ProtoMessage()    {}
func (*TablesAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4edb25480826735, []int{2}
}

func (m *TablesAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesAnnotation.Unmarshal(m, b)
}
func (m *TablesAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesAnnotation.Marshal(b, m, deterministic)
}
func (m *TablesAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesAnnotation.Merge(m, src)
}
func (m *TablesAnnotation) XXX_Size() int {
	return xxx_messageInfo_TablesAnnotation.Size(m)
}
func (m *TablesAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_TablesAnnotation proto.InternalMessageInfo

func (m *TablesAnnotation) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *TablesAnnotation) GetPredictionInterval() *DoubleRange {
	if m != nil {
		return m.PredictionInterval
	}
	return nil
}

func (m *TablesAnnotation) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TablesAnnotation) GetTablesModelColumnInfo() []*TablesModelColumnInfo {
	if m != nil {
		return m.TablesModelColumnInfo
	}
	return nil
}

func (m *TablesAnnotation) GetBaselineScore() float32 {
	if m != nil {
		return m.BaselineScore
	}
	return 0
}

// An information specific to given column and Tables Model, in context
// of the Model and the predictions created by it.
type TablesModelColumnInfo struct {
	// Output only. The name of the ColumnSpec describing the column. Not
	// populated when this proto is outputted to BigQuery.
	ColumnSpecName string `protobuf:"bytes,1,opt,name=column_spec_name,json=columnSpecName,proto3" json:"column_spec_name,omitempty"`
	// Output only. The display name of the column (same as the display_name of
	// its ColumnSpec).
	ColumnDisplayName string `protobuf:"bytes,2,opt,name=column_display_name,json=columnDisplayName,proto3" json:"column_display_name,omitempty"`
	// Output only. When given as part of a Model (always populated):
	// Measurement of how much model predictions correctness on the TEST data
	// depend on values in this column. A value between 0 and 1, higher means
	// higher influence. These values are normalized - for all input feature
	// columns of a given model they add to 1.
	//
	// When given back by Predict (populated iff
	// [feature_importance
	// param][google.cloud.automl.v1beta1.PredictRequest.params] is set) or Batch
	// Predict (populated iff
	// [feature_importance][google.cloud.automl.v1beta1.PredictRequest.params]
	// param is set):
	// Measurement of how impactful for the prediction returned for the given row
	// the value in this column was. Specifically, the feature importance
	// specifies the marginal contribution that the feature made to the prediction
	// score compared to the baseline score. These values are computed using the
	// Sampled Shapley method.
	FeatureImportance    float32  `protobuf:"fixed32,3,opt,name=feature_importance,json=featureImportance,proto3" json:"feature_importance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TablesModelColumnInfo) Reset()         { *m = TablesModelColumnInfo{} }
func (m *TablesModelColumnInfo) String() string { return proto.CompactTextString(m) }
func (*TablesModelColumnInfo) ProtoMessage()    {}
func (*TablesModelColumnInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4edb25480826735, []int{3}
}

func (m *TablesModelColumnInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesModelColumnInfo.Unmarshal(m, b)
}
func (m *TablesModelColumnInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesModelColumnInfo.Marshal(b, m, deterministic)
}
func (m *TablesModelColumnInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesModelColumnInfo.Merge(m, src)
}
func (m *TablesModelColumnInfo) XXX_Size() int {
	return xxx_messageInfo_TablesModelColumnInfo.Size(m)
}
func (m *TablesModelColumnInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesModelColumnInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TablesModelColumnInfo proto.InternalMessageInfo

func (m *TablesModelColumnInfo) GetColumnSpecName() string {
	if m != nil {
		return m.ColumnSpecName
	}
	return ""
}

func (m *TablesModelColumnInfo) GetColumnDisplayName() string {
	if m != nil {
		return m.ColumnDisplayName
	}
	return ""
}

func (m *TablesModelColumnInfo) GetFeatureImportance() float32 {
	if m != nil {
		return m.FeatureImportance
	}
	return 0
}

func init() {
	proto.RegisterType((*TablesDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TablesDatasetMetadata")
	proto.RegisterMapType((map[string]*CorrelationStats)(nil), "google.cloud.automl.v1beta1.TablesDatasetMetadata.TargetColumnCorrelationsEntry")
	proto.RegisterType((*TablesModelMetadata)(nil), "google.cloud.automl.v1beta1.TablesModelMetadata")
	proto.RegisterType((*TablesAnnotation)(nil), "google.cloud.automl.v1beta1.TablesAnnotation")
	proto.RegisterType((*TablesModelColumnInfo)(nil), "google.cloud.automl.v1beta1.TablesModelColumnInfo")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/tables.proto", fileDescriptor_a4edb25480826735)
}

var fileDescriptor_a4edb25480826735 = []byte{
	// 969 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc6, 0xbb, 0x4d, 0x0a, 0x13, 0x28, 0x89, 0xd3, 0x4d, 0x97, 0x0d, 0xa5, 0xab, 0x54, 0x80,
	0x55, 0x35, 0xde, 0x66, 0x01, 0x09, 0x2d, 0x5c, 0x90, 0x6c, 0x5a, 0x1a, 0x89, 0xb4, 0x91, 0x93,
	0x54, 0x02, 0x45, 0xb2, 0x66, 0xed, 0xb3, 0xee, 0xd0, 0xf1, 0x8c, 0x35, 0x33, 0x0e, 0xda, 0x5e,
	0xf1, 0x04, 0x3c, 0x03, 0x37, 0x3c, 0x03, 0x4f, 0xc0, 0x0d, 0x8f, 0xc2, 0x53, 0xa0, 0xf9, 0x71,
	0xb2, 0xbb, 0xd9, 0x98, 0x5c, 0x70, 0x67, 0xcf, 0xf7, 0xe3, 0xf1, 0x77, 0xce, 0x99, 0x41, 0x41,
	0xc6, 0x79, 0x46, 0xa1, 0x97, 0x50, 0x5e, 0xa6, 0x3d, 0x5c, 0x2a, 0x9e, 0xd3, 0xde, 0xf9, 0xce,
	0x08, 0x14, 0xde, 0xe9, 0x29, 0x3c, 0xa2, 0x20, 0xc3, 0x42, 0x70, 0xc5, 0xfd, 0x4d, 0xcb, 0x0c,
	0x0d, 0x33, 0xb4, 0xcc, 0xd0, 0x31, 0x3b, 0x4f, 0xea, 0x6c, 0x12, 0x8a, 0xa5, 0x24, 0x63, 0x92,
	0x60, 0x45, 0x38, 0xb3, 0x76, 0x9d, 0xed, 0x5a, 0x05, 0xa7, 0x65, 0xce, 0x62, 0x59, 0x40, 0xe2,
	0xe8, 0x8f, 0xeb, 0xe8, 0x29, 0x56, 0x38, 0x26, 0x0a, 0x72, 0x79, 0x63, 0xb6, 0x54, 0x58, 0x55,
	0xec, 0xda, 0x0c, 0x04, 0x66, 0x19, 0xdc, 0xc8, 0x57, 0x40, 0x26, 0x40, 0xca, 0xcb, 0x5f, 0x7c,
	0x54, 0x9b, 0x2d, 0xe4, 0x05, 0x17, 0x98, 0x3a, 0xee, 0xc7, 0x8e, 0x6b, 0xde, 0x46, 0xe5, 0xb8,
	0x27, 0x95, 0x28, 0x13, 0xe5, 0xd0, 0x07, 0xf3, 0xa8, 0x22, 0x39, 0x48, 0x85, 0xf3, 0x62, 0x4e,
	0x8e, 0x0b, 0xd2, 0xc3, 0x8c, 0x71, 0x65, 0xa2, 0x76, 0xdb, 0xde, 0xfa, 0xf3, 0x16, 0x6a, 0x9d,
	0x98, 0x5a, 0xee, 0x63, 0x85, 0x25, 0xa8, 0x43, 0x50, 0x58, 0xc7, 0xe0, 0xef, 0xa0, 0x56, 0x21,
	0x48, 0x8e, 0xc5, 0x24, 0x36, 0xc5, 0x36, 0x91, 0xc7, 0x24, 0x6d, 0x7b, 0x5d, 0x2f, 0x78, 0x2f,
	0xf2, 0x1d, 0x68, 0xc4, 0xc7, 0x05, 0x24, 0x07, 0xa9, 0x96, 0x28, 0x2c, 0x32, 0x50, 0xf1, 0x54,
	0x95, 0xb4, 0xa4, 0x61, 0x25, 0x16, 0x1c, 0x1a, 0xec, 0x52, 0xf2, 0x0b, 0x90, 0xec, 0xf5, 0x15,
	0x49, 0xd3, 0x4a, 0x2c, 0x38, 0x23, 0x79, 0x82, 0x5a, 0x39, 0x8d, 0x4b, 0x09, 0xf3, 0x92, 0x5b,
	0x46, 0xb2, 0x96, 0xd3, 0x53, 0x09, 0x33, 0x8a, 0xdf, 0x3c, 0xd4, 0x99, 0xdd, 0x58, 0xc2, 0x85,
	0x00, 0x6a, 0x93, 0x68, 0x2f, 0x77, 0x9b, 0xc1, 0x4a, 0xff, 0x28, 0xac, 0xe9, 0xe2, 0x70, 0x61,
	0x46, 0xe1, 0xc9, 0xd4, 0x0f, 0x0d, 0xa7, 0x2c, 0x9f, 0x32, 0x25, 0x26, 0x51, 0x5b, 0x5d, 0x03,
	0xfb, 0xcf, 0xd0, 0x9a, 0xe9, 0xb2, 0xb8, 0x2c, 0x52, 0xac, 0x20, 0xd6, 0x35, 0x6b, 0xdf, 0xee,
	0x7a, 0xc1, 0x4a, 0xbf, 0x53, 0x6d, 0xa3, 0x2a, 0x68, 0x78, 0x52, 0x15, 0x34, 0xfa, 0xd0, 0x88,
	0x4e, 0x8d, 0x46, 0xaf, 0x76, 0xde, 0xa2, 0xfb, 0xb5, 0x5b, 0xf0, 0x57, 0x51, 0xf3, 0x0d, 0x4c,
	0x5c, 0xc9, 0xf4, 0xa3, 0x3f, 0x44, 0x4b, 0xe7, 0x98, 0x96, 0x60, 0x6a, 0xb2, 0xd2, 0xdf, 0xae,
	0xfd, 0xeb, 0x29, 0xc3, 0x63, 0xfd, 0xe9, 0xc8, 0x6a, 0x07, 0x8d, 0xaf, 0xbd, 0xad, 0x5f, 0x97,
	0xd1, 0xba, 0x4d, 0xe5, 0x90, 0xa7, 0x40, 0x2f, 0xfa, 0xe6, 0x25, 0x7a, 0xc8, 0x0b, 0x45, 0x72,
	0xf2, 0xd6, 0xe8, 0x62, 0x3e, 0xfa, 0x19, 0x12, 0x45, 0xce, 0x21, 0x16, 0x90, 0x60, 0x4a, 0x63,
	0xfb, 0xf9, 0xb5, 0xae, 0x17, 0x34, 0x9e, 0xbf, 0x13, 0x75, 0xa7, 0xc9, 0x2f, 0x2b, 0x6e, 0x64,
	0xa8, 0xaf, 0x34, 0xd3, 0x3f, 0x45, 0x9f, 0x5d, 0x63, 0x58, 0x08, 0x48, 0x88, 0x9e, 0x2a, 0xe7,
	0xe9, 0x3b, 0xcf, 0x87, 0x0b, 0x3d, 0x8f, 0x2a, 0x76, 0x65, 0xeb, 0x5f, 0x6d, 0x56, 0x97, 0xca,
	0xe7, 0xff, 0x91, 0x4a, 0xd5, 0x5b, 0xd1, 0xea, 0x7c, 0x4b, 0xfb, 0x29, 0xea, 0x10, 0x56, 0x94,
	0x2a, 0x1e, 0x03, 0x56, 0xa5, 0x98, 0x69, 0x52, 0xd9, 0x6e, 0x9a, 0x56, 0xbb, 0xb1, 0xfd, 0x3d,
	0x63, 0xf5, 0xcc, 0x3a, 0x5d, 0xae, 0x4b, 0xff, 0x2b, 0xb4, 0xb1, 0x38, 0x13, 0x37, 0x04, 0xad,
	0x85, 0x09, 0xf8, 0x6f, 0x50, 0xdb, 0x1e, 0xdc, 0x71, 0xae, 0x6b, 0x56, 0xed, 0x8d, 0xb0, 0x31,
	0x6f, 0x2f, 0x99, 0xad, 0xf5, 0x6f, 0x30, 0x05, 0xa6, 0xde, 0x76, 0x37, 0x07, 0x6c, 0xcc, 0xa3,
	0x96, 0x5a, 0xb4, 0xec, 0x7f, 0x87, 0xee, 0x2b, 0x81, 0x09, 0x8b, 0x47, 0x65, 0xaa, 0x63, 0xce,
	0x09, 0xa5, 0x24, 0x66, 0x3c, 0x85, 0xf8, 0x35, 0x2f, 0x85, 0x9e, 0x3b, 0x2f, 0x68, 0x46, 0x1f,
	0x19, 0xd2, 0x9e, 0xe1, 0x1c, 0x6a, 0xca, 0x0b, 0x9e, 0xc2, 0x73, 0x4d, 0xf0, 0xbf, 0x45, 0x9b,
	0xd6, 0x21, 0xe1, 0x72, 0x81, 0xfe, 0xb6, 0xd1, 0xdf, 0x33, 0x94, 0x21, 0x97, 0xf3, 0xea, 0x2f,
	0xd1, 0x46, 0x4a, 0xa4, 0x39, 0xba, 0x00, 0x0b, 0x3a, 0x89, 0xa5, 0xe2, 0x45, 0x41, 0x58, 0xd6,
	0x7e, 0xbf, 0xeb, 0x05, 0xef, 0x46, 0x77, 0x1d, 0xfa, 0x54, 0x83, 0xc7, 0x0e, 0xdb, 0x7b, 0x84,
	0x02, 0x9c, 0xa6, 0x44, 0xe7, 0x86, 0x69, 0x7c, 0x4d, 0xe3, 0x25, 0x9c, 0x8d, 0x49, 0xb6, 0xf5,
	0x57, 0x03, 0xad, 0xda, 0x48, 0x76, 0x2f, 0x0e, 0x56, 0xff, 0x2e, 0x5a, 0x92, 0x09, 0x17, 0x60,
	0x86, 0xae, 0x11, 0xd9, 0x17, 0xff, 0x47, 0xb4, 0x5e, 0x08, 0x48, 0x49, 0x62, 0x9c, 0x08, 0x53,
	0x20, 0xce, 0x31, 0x35, 0xd5, 0x5a, 0xe9, 0x07, 0xb5, 0xa1, 0xef, 0xf3, 0x72, 0x44, 0x21, 0xd2,
	0x97, 0x8d, 0x3e, 0x75, 0x2b, 0x93, 0x03, 0xe7, 0xe1, 0x3f, 0x9e, 0x9d, 0xe8, 0x8d, 0x2b, 0x07,
	0x88, 0xe9, 0x77, 0x37, 0xba, 0xb5, 0x2d, 0xd0, 0xfc, 0xbf, 0x5b, 0xe0, 0x53, 0x74, 0x67, 0x84,
	0x25, 0x50, 0xc2, 0x20, 0xb6, 0xa1, 0x2c, 0x99, 0x50, 0x3e, 0xa8, 0x56, 0x8f, 0xf5, 0xe2, 0xd6,
	0xef, 0x5e, 0x75, 0x09, 0xcd, 0x1b, 0x04, 0x68, 0x75, 0xfa, 0x90, 0x67, 0x38, 0x07, 0x77, 0x98,
	0xdd, 0x49, 0x2e, 0xc6, 0xe1, 0x05, 0xce, 0xc1, 0x0f, 0xd1, 0xba, 0x63, 0xa6, 0x44, 0x16, 0x14,
	0x4f, 0x2c, 0xd9, 0xde, 0x3c, 0x6b, 0x16, 0xda, 0xb7, 0x88, 0xe1, 0x6f, 0x23, 0xbf, 0x9a, 0x50,
	0xa2, 0xaf, 0x5b, 0x85, 0x59, 0x02, 0xe6, 0xd6, 0x69, 0x44, 0x6b, 0x0e, 0x39, 0xb8, 0x00, 0xf6,
	0xfe, 0xf0, 0xd0, 0x83, 0x84, 0xe7, 0x75, 0xd1, 0x1c, 0x79, 0x3f, 0xed, 0x3a, 0x38, 0xe3, 0x14,
	0xb3, 0x2c, 0xe4, 0x22, 0xeb, 0x65, 0xc0, 0x4c, 0x21, 0x7a, 0x16, 0xc2, 0x05, 0x91, 0x0b, 0x6f,
	0xfd, 0x6f, 0xec, 0xeb, 0xdf, 0x8d, 0xcd, 0xef, 0x0d, 0xf1, 0x6c, 0xa8, 0x49, 0x67, 0xbb, 0xa5,
	0xe2, 0x87, 0xf4, 0xec, 0x95, 0x25, 0xfd, 0xd3, 0xf8, 0xc4, 0xa2, 0x83, 0x81, 0x81, 0x07, 0x03,
	0x83, 0xff, 0x30, 0x18, 0x38, 0xc2, 0x68, 0xd9, 0x7c, 0xec, 0x8b, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x4b, 0xb3, 0xa6, 0xbd, 0x09, 0x00, 0x00,
}
