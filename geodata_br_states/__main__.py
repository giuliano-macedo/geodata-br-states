from copy import deepcopy
import json
from typing import cast
import urllib.request
from pathlib import Path
from pyproj import Transformer

from geodata_br_states.types import Coordinate, GeoJson

_ROOT = Path(__file__).parent
_OUTPUT_DIR = _ROOT.parent.joinpath("geojson")


_URL = "https://geonode.paranagua.pr.gov.br/geoserver/ows?srsName=EPSG%3A31982&outputFormat=json&service=WFS&srs=EPSG%3A31982&request=GetFeature&typename=geonode%3Aa__031_003_estadosBrasil&version=1.0.0"


def get_geojson(url: str) -> GeoJson:
    with urllib.request.urlopen(url) as response:
        return json.loads(response.read().decode())


def fix_coordinate_projection(
    coordinates: Coordinate | list,
    transformer: Transformer,
) -> None:
    if not coordinates:
        return
    if isinstance(coordinates[0], float):
        coordinates = cast(Coordinate, coordinates)
        lng, lat = transformer.transform(coordinates[0], coordinates[1])
        coordinates[0], coordinates[1] = lat, lng
        return
    for coordinate in coordinates:
        coordinate = cast(list, coordinate)
        fix_coordinate_projection(coordinate, transformer)


def convert_all_coordiantes_from_geojson(
    geo_json: GeoJson,
    src_projection: str,
    dest_projection: str,
) -> GeoJson:
    converted_geo_json = deepcopy(geo_json)
    transformer = Transformer.from_crs(src_projection, dest_projection)
    for feature in converted_geo_json["features"]:
        for coordinate in feature["geometry"]["coordinates"]:
            fix_coordinate_projection(coordinate, transformer)
    return converted_geo_json


def path_as_markdown(path: Path) -> str:
    relative_path = path.relative_to(_ROOT.parent)
    return f"[{relative_path}]({relative_path})"


def write_individual_states(data: GeoJson) -> None:
    dir = _OUTPUT_DIR.joinpath("br_states")
    dir.mkdir(exist_ok=True)

    for feature in data["features"]:
        props = feature["properties"]

        file_out = dir.joinpath(f"br_{props['SIGLA'].lower()}.json")
        file_out.write_text(
            json.dumps(
                GeoJson(
                    features=[feature],
                    type=data["type"],
                ),
                ensure_ascii=False,
            )
        )
        print(f"* {props['Estado']}: {path_as_markdown(file_out)}")


def main() -> None:
    data = get_geojson(_URL)
    data = convert_all_coordiantes_from_geojson(data, "EPSG:31982", "EPSG:4326")
    data.pop("crs")
    data.pop("timeStamp")
    path_all_states = _OUTPUT_DIR.joinpath("br_states.json")
    path_all_states.write_text(json.dumps(data, ensure_ascii=False))
    print(f"all states: {path_as_markdown(path_all_states)}")
    write_individual_states(data)


if __name__ == "__main__":
    main()
