/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCourseclass,
    EntCourseclassFromJSON,
    EntCourseclassFromJSONTyped,
    EntCourseclassToJSON,
    EntDepartment,
    EntDepartmentFromJSON,
    EntDepartmentFromJSONTyped,
    EntDepartmentToJSON,
    EntInstructorRoom,
    EntInstructorRoomFromJSON,
    EntInstructorRoomFromJSONTyped,
    EntInstructorRoomToJSON,
    EntTitle,
    EntTitleFromJSON,
    EntTitleFromJSONTyped,
    EntTitleToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntInstructorInfoEdges
 */
export interface EntInstructorInfoEdges {
    /**
     * Courseclasses holds the value of the courseclasses edge.
     * @type {Array<EntCourseclass>}
     * @memberof EntInstructorInfoEdges
     */
    courseclasses?: Array<EntCourseclass>;
    /**
     * 
     * @type {EntDepartment}
     * @memberof EntInstructorInfoEdges
     */
    department?: EntDepartment;
    /**
     * 
     * @type {EntInstructorRoom}
     * @memberof EntInstructorInfoEdges
     */
    instructorroom?: EntInstructorRoom;
    /**
     * 
     * @type {EntTitle}
     * @memberof EntInstructorInfoEdges
     */
    title?: EntTitle;
}

export function EntInstructorInfoEdgesFromJSON(json: any): EntInstructorInfoEdges {
    return EntInstructorInfoEdgesFromJSONTyped(json, false);
}

export function EntInstructorInfoEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntInstructorInfoEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'courseclasses': !exists(json, 'courseclasses') ? undefined : ((json['courseclasses'] as Array<any>).map(EntCourseclassFromJSON)),
        'department': !exists(json, 'department') ? undefined : EntDepartmentFromJSON(json['department']),
        'instructorroom': !exists(json, 'instructorroom') ? undefined : EntInstructorRoomFromJSON(json['instructorroom']),
        'title': !exists(json, 'title') ? undefined : EntTitleFromJSON(json['title']),
    };
}

export function EntInstructorInfoEdgesToJSON(value?: EntInstructorInfoEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'courseclasses': value.courseclasses === undefined ? undefined : ((value.courseclasses as Array<any>).map(EntCourseclassToJSON)),
        'department': EntDepartmentToJSON(value.department),
        'instructorroom': EntInstructorRoomToJSON(value.instructorroom),
        'title': EntTitleToJSON(value.title),
    };
}


