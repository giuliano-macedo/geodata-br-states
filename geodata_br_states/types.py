from typing import Any, TypedDict


class Property(TypedDict):
    Estado: str
    SIGLA: str


Coordinate = list[float]


class Geometry(TypedDict):
    coordinates: list[list[list[Coordinate]]]


class Feature(TypedDict):
    properties: Property
    geometry: Geometry


class GeoJson(TypedDict, total=False):
    type: str
    features: list[Feature]
    crs: Any
    timeStamp: str
    totalFeatures: int
    numberMatched: int
    numberReturned: int