# GeoJson formatted Brazilian states perimiters

This project aims to distribute a GEOJson format of all brazilian states perimeters.

All states: [geojson/br_states.json](geojson/br_states.json)

Individual states:
* Acre: [geojson/br_states/br_ac.json](geojson/br_states/br_ac.json)
* Alagoas: [geojson/br_states/br_al.json](geojson/br_states/br_al.json)
* Amazonas: [geojson/br_states/br_am.json](geojson/br_states/br_am.json)
* Amapá: [geojson/br_states/br_ap.json](geojson/br_states/br_ap.json)
* Bahia: [geojson/br_states/br_ba.json](geojson/br_states/br_ba.json)
* Ceará: [geojson/br_states/br_ce.json](geojson/br_states/br_ce.json)
* Distrito Federal: [geojson/br_states/br_df.json](geojson/br_states/br_df.json)
* Espírito Santo: [geojson/br_states/br_es.json](geojson/br_states/br_es.json)
* Goiás: [geojson/br_states/br_go.json](geojson/br_states/br_go.json)
* Maranhão: [geojson/br_states/br_ma.json](geojson/br_states/br_ma.json)
* Minas Gerais: [geojson/br_states/br_mg.json](geojson/br_states/br_mg.json)
* Mato Grosso do Sul: [geojson/br_states/br_ms.json](geojson/br_states/br_ms.json)
* Mato Grosso: [geojson/br_states/br_mt.json](geojson/br_states/br_mt.json)
* Pará: [geojson/br_states/br_pa.json](geojson/br_states/br_pa.json)
* Paraíba: [geojson/br_states/br_pb.json](geojson/br_states/br_pb.json)
* Pernambuco: [geojson/br_states/br_pe.json](geojson/br_states/br_pe.json)
* Piauí: [geojson/br_states/br_pi.json](geojson/br_states/br_pi.json)
* Rio de Janeiro: [geojson/br_states/br_rj.json](geojson/br_states/br_rj.json)
* Rio Grande do Norte: [geojson/br_states/br_rn.json](geojson/br_states/br_rn.json)
* Rondônia: [geojson/br_states/br_ro.json](geojson/br_states/br_ro.json)
* Roraima: [geojson/br_states/br_rr.json](geojson/br_states/br_rr.json)
* Rio Grande do Sul: [geojson/br_states/br_rs.json](geojson/br_states/br_rs.json)
* Santa Catarina: [geojson/br_states/br_sc.json](geojson/br_states/br_sc.json)
* Sergipe: [geojson/br_states/br_se.json](geojson/br_states/br_se.json)
* Tocantins: [geojson/br_states/br_to.json](geojson/br_states/br_to.json)
* Paraná: [geojson/br_states/br_pr.json](geojson/br_states/br_pr.json)
* São Paulo: [geojson/br_states/br_sp.json](geojson/br_states/br_sp.json)


## Running the GeoJson generator

### Pre-requisites
* go 1.22

### Running
`go run main.go`

### Building
`go build`
then, the binary `geodata-br-states` will be created in the root directory

## Sources:
* Estados do Brasil - [LAGEAMB - UFPR](https://geonode.paranagua.pr.gov.br/groups/group/lageamb_ufpr/activity/):  https://geonode.paranagua.pr.gov.br/layers/geonode:a__031_003_estadosBrasil


## Related works
* GeoJSON Brazilian municipalities perimeters: https://github.com/tbrugz/geodata-br
