package entity

type DataTechRecord struct {
	tableName string `sql:"?db_schema.data_tech_records" json:"-"`

	*DataRecord
}

type TransitDataTechRecord struct {
	tableName string `sql:"?db_schema.data_tech_records" json:"-"`

	*TransitDataRecord
}
